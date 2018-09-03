package datamuse

import (
	"github.com/kostaspt/go-datamuse/datamuse/related"
	"github.com/kostaspt/go-datamuse/datamuse/vocabularies"
	"path"
	"strconv"
)

func (dm *Datamuse) Words() *Datamuse {
	dm.appendToPath("words")
	return dm
}

func (dm *Datamuse) Suggestions(s string) *Datamuse {
	dm.appendToPath("sug")
	dm.setQueryParam("s", s)
	return dm
}

func (dm *Datamuse) MeansLike(s string) *Datamuse {
	dm.setQueryParam("ml", s)
	return dm
}

func (dm *Datamuse) SpelledLike(s string) *Datamuse {
	dm.setQueryParam("sp", s)
	return dm
}

func (dm *Datamuse) SoundsLike(s string) *Datamuse {
	dm.setQueryParam("sl", s)
	return dm
}

func (dm *Datamuse) Related(code, s string) *Datamuse {
	dm.setQueryParam("rel_"+code, s)
	return dm
}

func (dm *Datamuse) RelatedNouns(s string) *Datamuse {
	return dm.Related(related.Nouns, s)
}

func (dm *Datamuse) RelatedAdjectives(s string) *Datamuse {
	return dm.Related(related.Adjectives, s)
}

func (dm *Datamuse) RelatedSynonyms(s string) *Datamuse {
	return dm.Related(related.Synonyms, s)
}

func (dm *Datamuse) RelatedTriggers(s string) *Datamuse {
	return dm.Related(related.Triggers, s)
}

func (dm *Datamuse) RelatedAntonyms(s string) *Datamuse {
	return dm.Related(related.Antonyms, s)
}

func (dm *Datamuse) RelatedKindOf(s string) *Datamuse {
	return dm.Related(related.KindOf, s)
}

func (dm *Datamuse) RelatedMoreGeneralThan(s string) *Datamuse {
	return dm.Related(related.MoreGeneralThan, s)
}

func (dm *Datamuse) RelatedComprises(s string) *Datamuse {
	return dm.Related(related.Comprises, s)
}

func (dm *Datamuse) RelatedPartOf(s string) *Datamuse {
	return dm.Related(related.PartOf, s)
}

func (dm *Datamuse) RelatedFrequentFollowers(s string) *Datamuse {
	return dm.Related(related.FrequentFollowers, s)
}

func (dm *Datamuse) RelatedFrequentPredecessors(s string) *Datamuse {
	return dm.Related(related.FrequentPredecessors, s)
}

func (dm *Datamuse) RelatedRhymes(s string) *Datamuse {
	return dm.Related(related.Rhymes, s)
}

func (dm *Datamuse) RelatedApproximateRhymes(s string) *Datamuse {
	return dm.Related(related.ApproximateRhymes, s)
}

func (dm *Datamuse) RelatedHomophones(s string) *Datamuse {
	return dm.Related(related.Homophones, s)
}

func (dm *Datamuse) RelatedConsonant(s string) *Datamuse {
	return dm.Related(related.Consonant, s)
}

func (dm *Datamuse) Vocabulary(s string) *Datamuse {
	dm.setQueryParam("v", s)
	return dm
}

func (dm *Datamuse) VocabularyDefault() *Datamuse {
	dm.setQueryParam("v", "")
	return dm
}

func (dm *Datamuse) VocabularySpanish() *Datamuse {
	return dm.Vocabulary(vocabularies.Spanish)
}

func (dm *Datamuse) VocabularyWikipedia() *Datamuse {
	return dm.Vocabulary(vocabularies.Wikipedia)
}

func (dm *Datamuse) Topics(s string) *Datamuse {
	dm.setQueryParam("topics", s)
	return dm
}

func (dm *Datamuse) LeftContext(s string) *Datamuse {
	dm.setQueryParam("lc", s)
	return dm
}

func (dm *Datamuse) RightContext(s string) *Datamuse {
	dm.setQueryParam("rc", s)
	return dm
}

func (dm *Datamuse) Max(n uint) *Datamuse {
	dm.setQueryParam("max", strconv.FormatUint(uint64(n), 10))
	return dm
}

func (dm *Datamuse) appendToPath(s string) {
	dm.ApiUrl.Path = path.Join(dm.ApiUrl.Path, s)
}

func (dm *Datamuse) setQueryParam(key, value string) {
	q := dm.ApiUrl.Query()

	q.Set(key, value)

	dm.ApiUrl.RawQuery = q.Encode()
}
