package datamuse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// In order to find words with a meaning similar to ringing in the ears.
func Test_Words_Means_Like(t *testing.T) {
	dm := New()
	res, err := dm.Words().MeansLike("ringing in the ears").Get()

	assert.NoError(t, err)
	assert.Equal(t, "https://api.datamuse.com/words?ml=ringing+in+the+ears", dm.URL())
	assert.Equal(t, "tinnitus", res[0].Word)
	assert.NotZero(t, res[0].Score)
	assert.EqualValues(t, []string{"syn", "n", "results_type:primary_external", "results_type:backfill_gloss"}, res[0].Tags)
}

// In order to find words related to duck that start with the letter b.
func Test_Words_Means_Like_Starts_With(t *testing.T) {
	dm := New()
	res, err := dm.Words().MeansLike("duck").SpelledLike("b*").Get()

	assert.NoError(t, err)
	assert.Equal(t, "https://api.datamuse.com/words?ml=duck&sp=b%2A", dm.URL())
	assert.Equal(t, "bird", res[0].Word)
	assert.NotZero(t, res[0].Score)
	assert.EqualValues(t, []string{"n", "results_type:primary_rel"}, res[0].Tags)
}

// In order to find words related to spoon that end with the letter a.
func Test_Words_Means_Like_Ends_With(t *testing.T) {
	dm := New()
	res, err := dm.Words().MeansLike("spoon").SpelledLike("*a").Get()

	assert.NoError(t, err)
	assert.Equal(t, "https://api.datamuse.com/words?ml=spoon&sp=%2Aa", dm.URL())
	assert.Equal(t, "spatula", res[0].Word)
	assert.NotZero(t, res[0].Score)
	assert.EqualValues(t, []string{"n", "results_type:primary_rel"}, res[0].Tags)
}

// In order to find words that sound like elefint.
func Test_Words_Sounds_Like(t *testing.T) {
	dm := New()
	res, err := dm.Words().SoundsLike("elefint").Get()

	assert.NoError(t, err)
	assert.Equal(t, "https://api.datamuse.com/words?sl=elefint", dm.URL())
	assert.Equal(t, "elefant", res[0].Word)
	assert.NotZero(t, res[0].Score)
	assert.EqualValues(t, 3, res[0].SyllablesCount)
}

// In order to find words that start with t, end in k, and have two letters in between.
func Test_Words_Spelled_Like_Pattern(t *testing.T) {
	dm := New()
	res, err := dm.Words().SpelledLike("t??k").Get()

	assert.NoError(t, err)
	assert.Equal(t, "https://api.datamuse.com/words?sp=t%3F%3Fk", dm.URL())
	assert.Equal(t, "talk", res[0].Word)
	assert.NotZero(t, res[0].Score)
}

// In order to find words that are spelled similarly to coneticut.
func Test_Words_Spelled_Like(t *testing.T) {
	dm := New()
	res, err := dm.Words().SpelledLike("coneticut").Get()

	assert.NoError(t, err)
	assert.Equal(t, "https://api.datamuse.com/words?sp=coneticut", dm.URL())
	assert.Equal(t, "conneticut", res[0].Word)
	assert.NotZero(t, res[0].Score)
}

// In order to find words that rhyme with forgetful.
func Test_Words_Related_Rhymes(t *testing.T) {
	dm := New()
	res, err := dm.Words().RelatedRhymes("forgetful").Get()

	assert.NoError(t, err)
	assert.Equal(t, "https://api.datamuse.com/words?rel_rhy=forgetful", dm.URL())
	assert.Equal(t, "fretful", res[0].Word)
	assert.NotZero(t, res[0].Score)
	assert.EqualValues(t, 2, res[0].SyllablesCount)
}

// In order to find words that rhyme with grape that are related to breakfast.
func Test_Words_Means_Like_Related_Rhymes(t *testing.T) {
	dm := New()
	res, err := dm.Words().MeansLike("breakfast").RelatedRhymes("grape").Get()

	assert.NoError(t, err)
	assert.Equal(t, "https://api.datamuse.com/words?ml=breakfast&rel_rhy=grape", dm.URL())
	assert.Equal(t, "crepe", res[0].Word)
	assert.NotZero(t, res[0].Score)
	assert.EqualValues(t, 1, res[0].SyllablesCount)
}

// In order to find adjectives that are often used to describe ocean.
func Test_Words_Related_Adjectives(t *testing.T) {
	dm := New()
	res, err := dm.Words().RelatedAdjectives("ocean").Get()

	assert.NoError(t, err)
	assert.Equal(t, "https://api.datamuse.com/words?rel_jjb=ocean", dm.URL())
	assert.Equal(t, "open", res[0].Word)
	assert.NotZero(t, res[0].Score)
}

// In order to find adjectives describing ocean sorted by how related they are to temperature.
func Test_Words_Related_Adjectives_By_Topics(t *testing.T) {
	dm := New()
	res, err := dm.Words().RelatedAdjectives("ocean").Topics("temperature").Get()

	assert.NoError(t, err)
	assert.Equal(t, "https://api.datamuse.com/words?rel_jjb=ocean&topics=temperature", dm.URL())
	assert.Equal(t, "cold", res[0].Word)
	assert.NotZero(t, res[0].Score)
}

// In order to find nouns that are often described by the adjective yellow.
func Test_Words_Related_Noun(t *testing.T) {
	dm := New()
	res, err := dm.Words().RelatedNouns("yellow").Get()

	assert.NoError(t, err)
	assert.Equal(t, "https://api.datamuse.com/words?rel_jja=yellow", dm.URL())
	assert.Equal(t, "fever", res[0].Word)
	assert.NotZero(t, res[0].Score)
}

// In order to find words that often follow "drink" in a sentence, that start with the letter w
func Test_Words_Left_Context(t *testing.T) {
	dm := New()
	res, err := dm.Words().LeftContext("drink").SpelledLike("w*").Get()

	assert.NoError(t, err)
	assert.Equal(t, "https://api.datamuse.com/words?lc=drink&sp=w%2A", dm.URL())
	assert.Equal(t, "with", res[0].Word)
	assert.NotZero(t, res[0].Score)
}

// In order to find words that are triggered by (strongly associated with) the word "cow"
func Test_Words_Related_Triggers(t *testing.T) {
	dm := New()
	res, err := dm.Words().RelatedTriggers("cow").Get()

	assert.NoError(t, err)
	assert.Equal(t, "https://api.datamuse.com/words?rel_trg=cow", dm.URL())
	assert.Equal(t, "dung", res[0].Word)
	assert.NotZero(t, res[0].Score)
}

// In order to find suggestions for the user if they have typed in rawand so far
func Test_Suggestions(t *testing.T) {
	dm := New()
	res, err := dm.Suggestions("rawand").Get()

	assert.NoError(t, err)
	assert.Equal(t, "https://api.datamuse.com/sug?s=rawand", dm.URL())
	assert.Equal(t, "rwanda", res[0].Word)
	assert.NotZero(t, res[0].Score)
}
