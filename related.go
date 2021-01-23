package datamuse

// Nouns - Popular nouns modified by the given adjective, per Google Books Ngrams
const RelatedNouns string = "jja"

// Adjectives - Popular adjectives used to modify the given noun, per Google Books Ngrams
const RelatedAdjectives string = "jjb"

// Synonyms (words contained within the same WordNet synset)
const RelatedSynonyms string = "syn"

// Triggers (words that are statistically associated with the query word in the same piece of text.)
const RelatedTriggers string = "trg"

// Antonyms (per WordNet)
const RelatedAntonyms string = "ant"

// KindOf (direct hypernyms, per WordNet)
const RelatedKindOf string = "spc"

// MoreGeneralThan (direct hyponyms, per WordNet)
const RelatedMoreGeneralThan string = "gen"

// Comprises (direct holonyms, per WordNet)
const RelatedComprises string = "com"

// PartOf (direct meronyms, per WordNet)
const RelatedPartOf string = "par"

// FrequentFollowers (w′ such that P(w′|w) ≥ 0.001, per Google Books Ngrams)
const RelatedFrequentFollowers string = "bga"

// FrequentPredecessors (w′ such that P(w|w′) ≥ 0.001, per Google Books Ngrams)
const RelatedFrequentPredecessors string = "bgb"

// Rhymes ("perfect" rhymes, per RhymeZone)
const RelatedRhymes string = "rhy"

// ApproximateRhymes (per RhymeZone)
const RelatedApproximateRhymes string = "nry"

// Homophones (sound-alike words)
const RelatedHomophones string = "hom"

// Consonant match
const RelatedConsonant string = "cns"
