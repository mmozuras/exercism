(ns bob
  (:require [clojure.string :refer (blank? lower-case upper-case)]))

(defn- no-lower-case-letters? [text]
  (= (upper-case text) text))

(defn- contains-upper-case-letters? [text]
  (not= (lower-case text) text))

(def shouting? (every-pred no-lower-case-letters? contains-upper-case-letters?))
(def silence? blank?)

(defn- question? [text]
  (= \? (last text)))

(defn response-for [text]
  (cond
    (shouting? text) "Woah, chill out!"
    (silence? text) "Fine. Be that way!"
    (question? text) "Sure."
    :else "Whatever."))
