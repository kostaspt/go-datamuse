package datamuse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatamuse_Words(t *testing.T) {
	assert.Equal(t, "https://api.datamuse.com/words", New().Words().URL())
}

func TestDatamuse_Suggestions(t *testing.T) {
	assert.Equal(t, "https://api.datamuse.com/sug?s=foo", New().Suggestions("foo").URL())
}

func TestDatamuse_MeansLike(t *testing.T) {
	expected := "https://api.datamuse.com/words?ml=foo"
	actual := New().Words().MeansLike("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_SpelledLike(t *testing.T) {
	expected := "https://api.datamuse.com/words?sp=foo"
	actual := New().Words().SpelledLike("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_SoundsLike(t *testing.T) {
	expected := "https://api.datamuse.com/words?sl=foo"
	actual := New().Words().SoundsLike("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_Related(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_jjb=foo"
	actual := New().Words().Related(RelatedAdjectives, "foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedNouns(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_jja=foo"
	actual := New().Words().RelatedNouns("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedAdjectives(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_jjb=foo"
	actual := New().Words().RelatedAdjectives("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedSynonyms(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_syn=foo"
	actual := New().Words().RelatedSynonyms("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedTriggers(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_trg=foo"
	actual := New().Words().RelatedTriggers("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedAntonyms(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_ant=foo"
	actual := New().Words().RelatedAntonyms("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedKindOf(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_spc=foo"
	actual := New().Words().RelatedKindOf("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedMoreGeneralThan(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_gen=foo"
	actual := New().Words().RelatedMoreGeneralThan("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedComprises(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_com=foo"
	actual := New().Words().RelatedComprises("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedPartOf(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_par=foo"
	actual := New().Words().RelatedPartOf("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedFrequentFollowers(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_bga=foo"
	actual := New().Words().RelatedFrequentFollowers("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedFrequentPredecessors(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_bgb=foo"
	actual := New().Words().RelatedFrequentPredecessors("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedRhymes(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_rhy=foo"
	actual := New().Words().RelatedRhymes("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedApproximateRhymes(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_nry=foo"
	actual := New().Words().RelatedApproximateRhymes("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedHomophones(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_hom=foo"
	actual := New().Words().RelatedHomophones("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedConsonant(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_cns=foo"
	actual := New().Words().RelatedConsonant("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_Vocabulary(t *testing.T) {
	expected := "https://api.datamuse.com/words?v=foo"
	actual := New().Words().Vocabulary("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_VocabularyDefault(t *testing.T) {
	expected := "https://api.datamuse.com/words?v="
	actual := New().Words().VocabularyDefault().URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_VocabularySpanish(t *testing.T) {
	expected := "https://api.datamuse.com/words?v=es"
	actual := New().Words().VocabularySpanish().URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_VocabularyWikipedia(t *testing.T) {
	expected := "https://api.datamuse.com/words?v=enwiki"
	actual := New().Words().VocabularyWikipedia().URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_Topics(t *testing.T) {
	expected := "https://api.datamuse.com/words?topics=foo"
	actual := New().Words().Topics("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_LeftContext(t *testing.T) {
	expected := "https://api.datamuse.com/words?lc=foo"
	actual := New().Words().LeftContext("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RightContext(t *testing.T) {
	expected := "https://api.datamuse.com/words?rc=foo"
	actual := New().Words().RightContext("foo").URL()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_Max(t *testing.T) {
	expected := "https://api.datamuse.com/words?max=42"
	actual := New().Words().Max(42).URL()

	assert.Equal(t, expected, actual)
}
