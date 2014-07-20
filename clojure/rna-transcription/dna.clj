(ns dna
  (:require [clojure.string :refer (join)]))

(def dna-to-rna {
  \C \C
  \G \G
  \A \A
  \T \U
})

(defn- valid-dna? [nucleotide]
  (contains? dna-to-rna nucleotide))

(defn- transcribe [nucleotide]
  (assert (valid-dna? nucleotide))
  (dna-to-rna nucleotide))

(defn to-rna [strand]
  (join "" (map transcribe strand)))
