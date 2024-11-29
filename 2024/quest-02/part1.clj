(defn parse-words [input]
  "WORDS:A,B,C,D -> ['A' 'B' 'C' 'C']"
  (as-> input $
        (str/split $ #"\n")
        (first $)
        (str/split $ #":")
        (second $)
        (str/split $ #",")
        (vec $)))

(defn parse-phrase [input]
  (last (str/split input #"\n")))

(defn part-1 [input]
  (let [words  (parse-words input)
        phrase (parse-phrase input)]
    (reduce + (map #(count (re-seq (re-pattern %) phrase)) words))))

(println 4  (part-1 (slurp "part1-sample.txt")))
(println 32 (part-1 (slurp "part1-actual.txt")))