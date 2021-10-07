package manager

import (
	"context"
	"fmt"
	"story_writer/src/constant"
	"story_writer/src/model"
)

func (mod *Module) AddWord(ctx context.Context, word string) (*model.AddWordResponse, error) {
	var sentenceId, paragraphId, storyId int64

	var storyResp model.AddWordResponse
	story, err := mod.useCases.StoryUseCase.GetStoryInProgress(ctx)
	if err != nil {
		return nil, err
	}

	paragraph, err := mod.useCases.ParagraphUseCase.GetParagraphInProgress(ctx)
	if err != nil {
		return nil, err
	}

	sentence, err := mod.useCases.SentenceUseCase.GetSentenceInProgress(ctx)
	if err != nil {
		return nil, err
	}

	// Improve it for more readability
	/* 	If story == nil => No story, paragraph, sentence is in progress. All have to be created.
	 	Else If paragraph == nil => No para, sentence is in progress. Only these two are required to be created.
		Else If sentence == nil => No Sentence is in progress, only sentence needs to be created.
		Else only word needs be inserted and attached to the current sentence.
	*/

	// This block is required to be created in transaction
	if story == nil {
		storyId, err = mod.useCases.StoryUseCase.InsertStory(ctx, &model.Story{Title: word, TitleLength: 1, Length: 0, Status: model.TitleInProgress})
		if err != nil {
			return nil, err
		}
		storyResp.Title = word
	} else {
		if story.Status == model.TitleInProgress && story.TitleLength < constant.MAX_TITLE_LENGTH {
			title := fmt.Sprintf("%s %s", story.Title, word)
			if story.TitleLength+1 == constant.MAX_TITLE_LENGTH {
				story.Status = model.TitleCompleted
			}

			storyId, err = mod.useCases.StoryUseCase.UpdateTitle(ctx, &model.Story{Title: title, TitleLength: story.TitleLength + 1, Status: story.Status, ID: story.ID})
			if err != nil {
				return nil, err
			}
			storyResp.Title = title
			storyResp.Id = storyId
			return &storyResp, nil
		}
		storyId = story.ID
	}

	storyResp.Id = storyId
	// No need to create paragraph, sentence, word till title is in progress
	if story == nil || story.Status == model.TitleInProgress {
		return &storyResp, nil
	}

	if paragraph == nil {
		paragraphId, err = mod.useCases.ParagraphUseCase.InsertParagraph(ctx, &model.Paragraph{StoryId: storyId, Length: 0, Status: model.InProgress})
		if err != nil {
			return nil, err
		}
	} else {
		paragraphId = paragraph.ID
	}

	if sentence == nil {
		sentenceId, err = mod.useCases.SentenceUseCase.InsertSentence(ctx, &model.Sentence{ParagraphId: paragraphId, Length: 0, Status: model.InProgress})
		if err != nil {
			return nil, err
		}
	} else {
		sentenceId = sentence.ID
	}

	_, err = mod.useCases.WordUseCase.InsertWord(ctx, &model.Word{Word: word, SentenceId: sentenceId})
	if err != nil {
		return nil, err
	}

	// update sentence word length
	wordsLength := 1
	if sentence != nil {
		wordsLength = sentence.Length + 1
	}

	_, err = mod.useCases.SentenceUseCase.IncrementLength(ctx, &model.Sentence{ID: sentenceId, Length: wordsLength, Status: model.InProgress})
	if err != nil {
		return nil, err
	}

	sentencesLength := 1
	//if words length has reached max, it means current sentence is Completed, so need to increase current paragraph length
	if wordsLength == constant.MAX_WORDS_LENGTH {
		if paragraph != nil {
			sentencesLength = paragraph.Length + 1
		}

		_, err = mod.useCases.ParagraphUseCase.IncrementLength(ctx, &model.Paragraph{ID: paragraphId, Length: sentencesLength, Status: model.InProgress})
		if err != nil {
			return nil, err
		}
	}

	paraLength := 1
	//if sentence length has reached max, it means current paragraph is Completed, so need to increase current story length
	if sentencesLength == constant.MAX_SENTENCE_LENGTH {
		if story != nil {
			paraLength = story.Length + 1
		}

		_, err = mod.useCases.StoryUseCase.IncrementLength(ctx, &model.Story{ID: storyId, Length: paraLength, Status: model.InProgress})
		if err != nil {
			return nil, err
		}
	}

	// txn block ends
	storyResp.CurrentSentence, err = mod.prepareSentence(ctx, sentenceId, storyResp.CurrentSentence)
	if err != nil {
		return nil, err
	}

	if story != nil {
		storyResp.Title = story.Title
	}

	return &storyResp, nil
}

func (mod *Module) GetStories(ctx context.Context, params model.Params) (model.StoriesResponse, error) {
	var response model.StoriesResponse

	stories, err := mod.useCases.StoryUseCase.GetStories(ctx, params)
	if err != nil {
		return response, err
	}

	response = model.StoriesResponse{Limit: params.Limit, Offset: params.Offset, Count: len(stories), Results: stories}

	return response, nil
}

func (mod *Module) GetStoryById(ctx context.Context, storyId int64) (*model.StoryResponse, error) {
	storyResp, err := mod.useCases.StoryUseCase.GetStoryById(ctx, storyId)
	if err != nil {
		return nil, err
	}

	sentenceParaIds, err := mod.useCases.SentenceUseCase.GetSentenceParaIdsByStoryId(ctx, storyId)
	if err != nil {
		return nil, err
	}

	paragraphMap := make(map[int64][]string)
	for _, ids := range sentenceParaIds {
		sentenceStr, err := mod.prepareSentence(ctx, ids.SentenceId, "")
		if err != nil {
			continue
		}
		paragraphMap[ids.ParagraphId] = append(paragraphMap[ids.ParagraphId], sentenceStr)
	}

	for _, sentences := range paragraphMap {
		storyResp.Paragraphs = append(storyResp.Paragraphs, model.ParagraphResponse{Sentences: sentences})
	}

	return storyResp, nil
}

func (mod *Module) prepareSentence(ctx context.Context, sentenceId int64, currentSentence string) (string, error) {
	words, err := mod.useCases.WordUseCase.GetWordsBySentenceId(ctx, sentenceId)
	if err != nil {
		return currentSentence, err
	}

	for _, wrd := range words {
		if currentSentence == "" {
			currentSentence = wrd.Word
		} else {
			currentSentence = fmt.Sprintf("%s %s", currentSentence, wrd.Word)
		}
	}
	return currentSentence, nil
}
