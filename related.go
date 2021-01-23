package datamuse

// RelatedNouns - Popular nouns modified by the given adjective, per Google Books Ngrams
const RelatedNouns string = "jja"

// RelatedAdjectives - Popular adjectives used to modify the given noun, per Google Books Ngrams
const RelatedAdjectives string = "jjb"

// RelatedSynonyms (words contained within the same WordNet synset)
const RelatedSynonyms string = "syn"

// RelatedTriggers (words that are statistically associated with the query word in the same piece of text.)
const RelatedTriggers string = "trg"

// RelatedAntonyms (per WordNet)
const RelatedAntonyms string = "ant"

// RelatedKindOf (direct hypernyms, per WordNet)
const RelatedKindOf string = "spc"

// RelatedMoreGeneralThan (direct hyponyms, per WordNet)
const RelatedMoreGeneralThan string = "gen"

// RelatedComprises (direct holonyms, per WordNet)
const RelatedComprises string = "com"

// RelatedPartOf (direct meronyms, per WordNet)
const RelatedPartOf string = "par"

// RelatedFrequentFollowers (w′ such that P(w′|w) ≥ 0.001, per Google Books Ngrams)
const RelatedFrequentFollowers string = "bga"

// RelatedFrequentPredecessors (w′ such that P(w|w′) ≥ 0.001, per Google Books Ngrams)
const RelatedFrequentPredecessors string = "bgb"

// RelatedRhymes ("perfect" rhymes, per RhymeZone)
const RelatedRhymes string = "rhy"

// RelatedApproximateRhymes (per RhymeZone)
const RelatedApproximateRhymes string = "nry"

// RelatedHomophones (sound-alike words)
const RelatedHomophones string = "hom"

// RelatedConsonant match
const RelatedConsonant string = "cns"
