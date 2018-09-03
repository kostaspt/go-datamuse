package datamuse

import (
	"github.com/kostaspt/go-datamuse/datamuse/related"
	"github.com/kostaspt/go-datamuse/datamuse/vocabularies"
	"path"
	"strconv"
)

// Words endpoint returns a list of words (and multiword expressions) from a given vocabulary that match a given set of
// constraints.
func (dm *Datamuse) Words() *Datamuse {
	dm.appendToPath("words")
	return dm
}

// Suggestions endpoint provides word suggestions given a partially-entered query.
func (dm *Datamuse) Suggestions(s string) *Datamuse {
	dm.appendToPath("sug")
	dm.setQueryParam("s", s)
	return dm
}

// MeansLike - require that the results have a meaning related to this string value, which can be any word or
// sequence of words. (This is effectively the reverse dictionary feature of OneLook.)
func (dm *Datamuse) MeansLike(s string) *Datamuse {
	dm.setQueryParam("ml", s)
	return dm
}

// SoundsLike - require that the results are pronounced similarly to this string of characters. (If the string
// of characters doesn't have a known pronunciation, the system will make its best guess using a text-to-phonemes
// algorithm.)
func (dm *Datamuse) SoundsLike(s string) *Datamuse {
	dm.setQueryParam("sl", s)
	return dm
}

// SpelledLike - require that the results are spelled similarly to this string of characters, or that they
// match this wildcard pattern. A pattern can include any combination of alphanumeric characters, spaces, and two
// reserved characters that represent placeholders — * (which matches any number of characters) and ? (which matches
// exactly one character).
func (dm *Datamuse) SpelledLike(s string) *Datamuse {
	dm.setQueryParam("sp", s)
	return dm
}

// Related - require that the results, when paired with the word in this parameter, are in a predefined
// lexical relation indicated by [code]. Any number of these parameters may be specified any number of times. An
// assortment of semantic, phonetic, and corpus-statistics-based relations are available. At this time, these relations
// are available for English-language vocabularies only.
func (dm *Datamuse) Related(code, s string) *Datamuse {
	dm.setQueryParam("rel_"+code, s)
	return dm
}

// RelatedNouns - Popular nouns modified by the given adjective, per Google Books Ngrams.
func (dm *Datamuse) RelatedNouns(s string) *Datamuse {
	return dm.Related(related.Nouns, s)
}

// RelatedAdjectives - Popular adjectives used to modify the given noun, per Google Books Ngrams.
func (dm *Datamuse) RelatedAdjectives(s string) *Datamuse {
	return dm.Related(related.Adjectives, s)
}

// RelatedSynonyms - Synonyms (words contained within the same WordNet synset).
func (dm *Datamuse) RelatedSynonyms(s string) *Datamuse {
	return dm.Related(related.Synonyms, s)
}

// RelatedTriggers - "Triggers" (words that are statistically associated with the query word in the same piece
// of text.)
func (dm *Datamuse) RelatedTriggers(s string) *Datamuse {
	return dm.Related(related.Triggers, s)
}

// RelatedAntonyms - Antonyms (per WordNet)
func (dm *Datamuse) RelatedAntonyms(s string) *Datamuse {
	return dm.Related(related.Antonyms, s)
}

// RelatedKindOf - "Kind of" (direct hypernyms, per WordNet)
func (dm *Datamuse) RelatedKindOf(s string) *Datamuse {
	return dm.Related(related.KindOf, s)
}

// RelatedMoreGeneralThan - "More general than" (direct hyponyms, per WordNet)
func (dm *Datamuse) RelatedMoreGeneralThan(s string) *Datamuse {
	return dm.Related(related.MoreGeneralThan, s)
}

// RelatedComprises - "Comprises" (direct holonyms, per WordNet)
func (dm *Datamuse) RelatedComprises(s string) *Datamuse {
	return dm.Related(related.Comprises, s)
}

// RelatedPartOf - "Part of" (direct meronyms, per WordNet)
func (dm *Datamuse) RelatedPartOf(s string) *Datamuse {
	return dm.Related(related.PartOf, s)
}

// RelatedFrequentFollowers - Frequent followers (w′ such that P(w′|w) ≥ 0.001, per Google Books Ngrams)
func (dm *Datamuse) RelatedFrequentFollowers(s string) *Datamuse {
	return dm.Related(related.FrequentFollowers, s)
}

// RelatedFrequentPredecessors - Frequent predecessors (w′ such that P(w|w′) ≥ 0.001, per Google Books Ngrams)
func (dm *Datamuse) RelatedFrequentPredecessors(s string) *Datamuse {
	return dm.Related(related.FrequentPredecessors, s)
}

// RelatedRhymes - Rhymes ("perfect" rhymes, per RhymeZone)
func (dm *Datamuse) RelatedRhymes(s string) *Datamuse {
	return dm.Related(related.Rhymes, s)
}

// RelatedApproximateRhymes - Approximate rhymes (per RhymeZone)
func (dm *Datamuse) RelatedApproximateRhymes(s string) *Datamuse {
	return dm.Related(related.ApproximateRhymes, s)
}

// RelatedHomophones - Homophones (sound-alike words)
func (dm *Datamuse) RelatedHomophones(s string) *Datamuse {
	return dm.Related(related.Homophones, s)
}

// RelatedConsonant - Consonant match
func (dm *Datamuse) RelatedConsonant(s string) *Datamuse {
	return dm.Related(related.Consonant, s)
}

// Vocabulary - Identifier for the vocabulary to use.
func (dm *Datamuse) Vocabulary(s string) *Datamuse {
	dm.setQueryParam("v", s)
	return dm
}

// VocabularyDefault - The default 550,000-term vocabulary of English words and multiword expressions is used.
func (dm *Datamuse) VocabularyDefault() *Datamuse {
	dm.setQueryParam("v", "")
	return dm
}

// VocabularySpanish - The 500,000-term vocabulary of words from Spanish-language books is used.
func (dm *Datamuse) VocabularySpanish() *Datamuse {
	return dm.Vocabulary(vocabularies.Spanish)
}

// VocabularyWikipedia - The approximately 6 million-term vocabulary of article titles from the English-language
// Wikipedia is used (updated monthly).
func (dm *Datamuse) VocabularyWikipedia() *Datamuse {
	return dm.Vocabulary(vocabularies.Wikipedia)
}

// Topics - An optional hint to the system about the theme of the document being written. Results will be skewed toward
// these topics. At most 5 words can be specified. Space or comma delimited. Nouns work best.
func (dm *Datamuse) Topics(s string) *Datamuse {
	dm.setQueryParam("topics", s)
	return dm
}

// LeftContext - An optional hint to the system about the word that appears immediately to the left of the target word
// in a sentence. (At this time, only a single word may be specified.)
func (dm *Datamuse) LeftContext(s string) *Datamuse {
	dm.setQueryParam("lc", s)
	return dm
}

// RightContext - An optional hint to the system about the word that appears immediately to the right of the target word
// in a sentence. (At this time, only a single word may be specified.)
func (dm *Datamuse) RightContext(s string) *Datamuse {
	dm.setQueryParam("rc", s)
	return dm
}

// Max - Maximum number of results to return, not to exceed 1000. (default: 100)
func (dm *Datamuse) Max(n uint) *Datamuse {
	dm.setQueryParam("max", strconv.FormatUint(uint64(n), 10))
	return dm
}

func (dm *Datamuse) appendToPath(s string) {
	dm.APIURL.Path = path.Join(dm.APIURL.Path, s)
}

func (dm *Datamuse) setQueryParam(key, value string) {
	q := dm.APIURL.Query()

	q.Set(key, value)

	dm.APIURL.RawQuery = q.Encode()
}
