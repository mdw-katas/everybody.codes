(def potion-costs {\A 0 \B 1 \C 3})
(defn potion-cost [input] (reduce + (map potion-costs input)))
(println 5    (potion-cost (slurp "part1-sample.txt")))
(println 1395 (potion-cost (slurp "part1-actual.txt")))
