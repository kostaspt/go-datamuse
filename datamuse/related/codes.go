package related

// Nouns - Popular nouns modified by the given adjective, per Google Books Ngrams
const Nouns string = "jja"

// Adjectives - Popular adjectives used to modify the given noun, per Google Books Ngrams
const Adjectives string = "jjb"

// Synonyms (words contained within the same WordNet synset)
const Synonyms string = "syn"

// Triggers (words that are statistically associated with the query word in the same piece of text.)
const Triggers string = "trg"

// Antonyms (per WordNet)
const Antonyms string = "ant"

// KindOf (direct hypernyms, per WordNet)
const KindOf string = "spc"

// MoreGeneralThan (direct hyponyms, per WordNet)
const MoreGeneralThan string = "gen"

// Comprises (direct holonyms, per WordNet)
const Comprises string = "com"

// PartOf (direct meronyms, per WordNet)
const PartOf string = "par"

// FrequentFollowers (w′ such that P(w′|w) ≥ 0.001, per Google Books Ngrams)
const FrequentFollowers string = "bga"

// FrequentPredecessors (w′ such that P(w|w′) ≥ 0.001, per Google Books Ngrams)
const FrequentPredecessors string = "bgb"

// Rhymes ("perfect" rhymes, per RhymeZone)
const Rhymes string = "rhy"

// ApproximateRhymes (per RhymeZone)
const ApproximateRhymes string = "nry"

// Homophones (sound-alike words)
const Homophones string = "hom"

// Consonant match
const Consonant string = "cns"
