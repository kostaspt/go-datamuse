package related

// Popular nouns modified by the given adjective, per Google Books Ngrams
const Nouns string = "jja"

// Popular adjectives used to modify the given noun, per Google Books Ngrams
const Adjectives string = "jjb"

// Synonyms (words contained within the same WordNet synset)
const Synonyms string = "syn"

// "Triggers" (words that are statistically associated with the query word in the same piece of text.)
const Triggers string = "trg"

// Antonyms (per WordNet)
const Antonyms string = "ant"

// "Kind of" (direct hypernyms, per WordNet)
const KindOf string = "spc"

// "More general than" (direct hyponyms, per WordNet)
const MoreGeneralThan string = "gen"

// "Comprises" (direct holonyms, per WordNet)
const Comprises string = "com"

// "Part of" (direct meronyms, per WordNet)
const PartOf string = "par"

// Frequent followers (w′ such that P(w′|w) ≥ 0.001, per Google Books Ngrams)
const FrequentFollowers string = "bga"

// Frequent predecessors (w′ such that P(w|w′) ≥ 0.001, per Google Books Ngrams)
const FrequentPredecessors string = "bgb"

// Rhymes ("perfect" rhymes, per RhymeZone)
const Rhymes string = "rhy"

// Approximate rhymes (per RhymeZone)
const ApproximateRhymes string = "nry"

// Homophones (sound-alike words)
const Homophones string = "hom"

// Consonant match
const Consonant string = "cns"
