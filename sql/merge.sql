SELECT QuestionsTags.question_id, Questions.title, Questions.slug, Questions.difficulty, string_agg(Tags.name, ', ')
FROM QuestionsTags
LEFT JOIN question ON Questions.id = QuestionsTags.question_id
LEFT JOIN tag ON Tags.id = QuestionsTags.tag_id
GROUP BY QuestionsTags.question_id, Questions.title, Questions.slug, Questions.difficulty;