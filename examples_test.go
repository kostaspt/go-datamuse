package datamuse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// In order to find words with a meaning similar to ringing in the ears.
// Endpoint: https://api.datamuse.com/words?ml=ringing+in+the+ears
func Test_Words_Means_Like(t *testing.T) {
	res, err := New().Words().MeansLike("ringing in the ears").Get()

	assert.NoError(t, err)
	assert.Equal(t, "tinnitus", res[0].Word)
	assert.NotZero(t, res[0].Score)
	assert.EqualValues(t, []string{"syn", "n"}, res[0].Tags)
}

// In order to find words related to duck that start with the letter b.
// Endpoint: https://api.datamuse.com/words?ml=duck&sp=b*
func Test_Words_Means_Like_Starts_With(t *testing.T) {
	res, err := New().Words().MeansLike("duck").SpelledLike("b*").Get()

	assert.NoError(t, err)
	assert.Equal(t, "bird", res[0].Word)
	assert.NotZero(t, res[0].Score)
	assert.EqualValues(t, []string{"n"}, res[0].Tags)
}

// In order to find words related to spoon that end with the letter a.
// Endpoint: https://api.datamuse.com/words?ml=spoon&sp=*a
func Test_Words_Means_Like_Ends_With(t *testing.T) {
	res, err := New().Words().MeansLike("spoon").SpelledLike("*a").Get()

	assert.NoError(t, err)
	assert.Equal(t, "spatula", res[0].Word)
	assert.NotZero(t, res[0].Score)
	assert.EqualValues(t, []string{"n"}, res[0].Tags)
}

// In order to find words that sound like elefint.
// Endpoint: https://api.datamuse.com/words?sl=elefint
func Test_Words_Sounds_Like(t *testing.T) {
	res, err := New().Words().SoundsLike("elefint").Get()

	assert.NoError(t, err)
	assert.Equal(t, "elephant", res[0].Word)
	assert.NotZero(t, res[0].Score)
	assert.EqualValues(t, 3, res[0].SyllablesCount)
}

// In order to find words that start with t, end in k, and have two letters in between.
// Endpoint: https://api.datamuse.com/words?sp=t??k
func Test_Words_Spelled_Like_Pattern(t *testing.T) {
	res, err := New().Words().SpelledLike("t??k").Get()

	assert.NoError(t, err)
	assert.Equal(t, "talk", res[0].Word)
	assert.NotZero(t, res[0].Score)
}

// In order to find words that are spelled similarly to coneticut.
// Endpoint: https://api.datamuse.com/words?sp=coneticut
func Test_Words_Spelled_Like(t *testing.T) {
	res, err := New().Words().SpelledLike("coneticut").Get()

	assert.NoError(t, err)
	assert.Equal(t, "conneticut", res[0].Word)
	assert.NotZero(t, res[0].Score)
}

// In order to find words that rhyme with forgetful.
// Endpoint: https://api.datamuse.com/words?rel_rhy=forgetful
func Test_Words_Related_Rhymes(t *testing.T) {
	res, err := New().Words().RelatedRhymes("forgetful").Get()

	assert.NoError(t, err)
	assert.Equal(t, "fretful", res[0].Word)
	assert.NotZero(t, res[0].Score)
	assert.EqualValues(t, 2, res[0].SyllablesCount)
}

// In order to find words that rhyme with grape that are related to breakfast.
// Endpoint: https://api.datamuse.com/words?ml=breakfast&rel_rhy=grape
func Test_Words_Means_Like_Related_Rhymes(t *testing.T) {
	res, err := New().Words().MeansLike("breakfast").RelatedRhymes("grape").Get()

	assert.NoError(t, err)
	assert.Equal(t, "crepe", res[0].Word)
	assert.NotZero(t, res[0].Score)
	assert.EqualValues(t, 1, res[0].SyllablesCount)
}

// In order to find adjectives that are often used to describe ocean.
// Endpoint: https://api.datamuse.com/words?rel_jjb=ocean
func Test_Words_Related_Adjectives(t *testing.T) {
	res, err := New().Words().RelatedAdjectives("ocean").Get()

	assert.NoError(t, err)
	assert.Equal(t, "open", res[0].Word)
	assert.NotZero(t, res[0].Score)
}

// In order to find adjectives describing ocean sorted by how related they are to temperature.
// Endpoint: https://api.datamuse.com/words?rel_jjb=ocean&topics=temperature
func Test_Words_Related_Adjectives_By_Topics(t *testing.T) {
	res, err := New().Words().RelatedAdjectives("ocean").Topics("temperature").Get()

	assert.NoError(t, err)
	assert.Equal(t, "cold", res[0].Word)
	assert.NotZero(t, res[0].Score)
}

// In order to find nouns that are often described by the adjective yellow.
// Endpoint: https://api.datamuse.com/words?rel_jja=yellow
func Test_Words_Related_Noun(t *testing.T) {
	res, err := New().Words().RelatedNouns("yellow").Get()

	assert.NoError(t, err)
	assert.Equal(t, "fever", res[0].Word)
	assert.NotZero(t, res[0].Score)
}

// In order to find words that often follow "drink" in a sentence, that start with the letter w
// Endpoint: https://api.datamuse.com/words?lc=drink&sp=w*
func Test_Words_Left_Context(t *testing.T) {
	res, err := New().Words().LeftContext("yellow").SpelledLike("w*").Get()

	assert.NoError(t, err)
	assert.Equal(t, "with", res[0].Word)
	assert.NotZero(t, res[0].Score)
}

// In order to find words that are triggered by (strongly associated with) the word "cow"
// Endpoint: https://api.datamuse.com/words?rel_trg=cow
func Test_Words_Related_Triggers(t *testing.T) {
	res, err := New().Words().RelatedTriggers("cow").Get()

	assert.NoError(t, err)
	assert.Equal(t, "dung", res[0].Word)
	assert.NotZero(t, res[0].Score)
}

// In order to find suggestions for the user if they have typed in rawand so far
// Endpoint: https://api.datamuse.com/sug?s=rawand
func Test_Suggestions(t *testing.T) {
	res, err := New().Suggestions("rawand").Get()

	assert.NoError(t, err)
	assert.Equal(t, "rwanda", res[0].Word)
	assert.NotZero(t, res[0].Score)
}
