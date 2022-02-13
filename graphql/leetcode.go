package graphql

import (
	"context"

	"github.com/machinebox/graphql"
	"github.com/weehong/leetcode-tracker/model"
)

func Query() (*model.LeetCode, error) {
	client := graphql.NewClient("https://leetcode.com/graphql")
	request := graphql.NewRequest(`
		query problemsetQuestionList(
			$categorySlug: String
			$limit: Int
			$skip: Int
			$filters: QuestionListFilterInput
		) {
			problemsetQuestionList: questionList(
				categorySlug: $categorySlug
				limit: $limit
				skip: $skip
				filters: $filters
			) {
				total: totalNum
				questions: data {
					acRate
					difficulty
					freqBar
					frontendQuestionId: questionFrontendId
					isFavor
					paidOnly: isPaidOnly
					status
					title
					titleSlug
					topicTags {
						name
						id
						slug
					}
					hasSolution
					hasVideoSolution
				}
			}
		}
	`)

	request.Var("categorySlug", "")
	request.Var("skip", 0)
	request.Var("limit", -1)
	request.Var("filters", map[string]string{})

	ctx := context.Background()
	var response model.LeetCode

	if err := client.Run(ctx, request, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
