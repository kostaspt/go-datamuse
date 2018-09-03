package datamuse

import (
	"github.com/kostaspt/go-datamuse/datamuse/related"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDatamuse_Words(t *testing.T) {
	assert.Equal(t, "https://api.datamuse.com/words", New().Words().Hyperlink())
}

func TestDatamuse_Suggestions(t *testing.T) {
	assert.Equal(t, "https://api.datamuse.com/sug?s=foo", New().Suggestions("foo").Hyperlink())
}

func TestDatamuse_MeansLike(t *testing.T) {
	expected := "https://api.datamuse.com/words?ml=foo"
	actual := New().Words().MeansLike("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_SpelledLike(t *testing.T) {
	expected := "https://api.datamuse.com/words?sp=foo"
	actual := New().Words().SpelledLike("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_SoundsLike(t *testing.T) {
	expected := "https://api.datamuse.com/words?sl=foo"
	actual := New().Words().SoundsLike("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_Related(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_jjb=foo"
	actual := New().Words().Related(related.Adjectives, "foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedNouns(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_jja=foo"
	actual := New().Words().RelatedNouns("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedAdjectives(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_jjb=foo"
	actual := New().Words().RelatedAdjectives("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedSynonyms(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_syn=foo"
	actual := New().Words().RelatedSynonyms("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedTriggers(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_trg=foo"
	actual := New().Words().RelatedTriggers("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedAntonyms(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_ant=foo"
	actual := New().Words().RelatedAntonyms("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedKindOf(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_spc=foo"
	actual := New().Words().RelatedKindOf("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedMoreGeneralThan(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_gen=foo"
	actual := New().Words().RelatedMoreGeneralThan("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedComprises(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_com=foo"
	actual := New().Words().RelatedComprises("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedPartOf(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_par=foo"
	actual := New().Words().RelatedPartOf("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedFrequentFollowers(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_bga=foo"
	actual := New().Words().RelatedFrequentFollowers("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedFrequentPredecessors(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_bgb=foo"
	actual := New().Words().RelatedFrequentPredecessors("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedRhymes(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_rhy=foo"
	actual := New().Words().RelatedRhymes("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedApproximateRhymes(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_nry=foo"
	actual := New().Words().RelatedApproximateRhymes("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedHomophones(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_hom=foo"
	actual := New().Words().RelatedHomophones("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RelatedConsonant(t *testing.T) {
	expected := "https://api.datamuse.com/words?rel_cns=foo"
	actual := New().Words().RelatedConsonant("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_Vocabulary(t *testing.T) {
	expected := "https://api.datamuse.com/words?v=foo"
	actual := New().Words().Vocabulary("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_VocabularyDefault(t *testing.T) {
	expected := "https://api.datamuse.com/words?v="
	actual := New().Words().VocabularyDefault().Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_VocabularySpanish(t *testing.T) {
	expected := "https://api.datamuse.com/words?v=es"
	actual := New().Words().VocabularySpanish().Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_VocabularyWikipedia(t *testing.T) {
	expected := "https://api.datamuse.com/words?v=enwiki"
	actual := New().Words().VocabularyWikipedia().Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_Topics(t *testing.T) {
	expected := "https://api.datamuse.com/words?topics=foo"
	actual := New().Words().Topics("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_LeftContext(t *testing.T) {
	expected := "https://api.datamuse.com/words?lc=foo"
	actual := New().Words().LeftContext("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_RightContext(t *testing.T) {
	expected := "https://api.datamuse.com/words?rc=foo"
	actual := New().Words().RightContext("foo").Hyperlink()

	assert.Equal(t, expected, actual)
}

func TestDatamuse_Max(t *testing.T) {
	expected := "https://api.datamuse.com/words?max=42"
	actual := New().Words().Max(42).Hyperlink()

	assert.Equal(t, expected, actual)
}
