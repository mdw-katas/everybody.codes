(def potion-costs {\x -1 \A (inc 0) \B (inc 1) \C (inc 3) \D (inc 5)})
(defn potion-cost [input] (reduce + (map potion-costs input)))
(defn part2 [filename]
  (->> (slurp filename)
    (partition 2)
    (remove #(= % '(\x \x))) ; remove double-x pairs
    (apply concat)
    (apply str)
    potion-cost))

(println 0    (potion-cost "Ax"))
(println 6    (potion-cost "BC"))
(println 12   (potion-cost "DD"))
(println 5    (potion-cost "CA"))
(println 5    (potion-cost "xD"))

(println 28   (part2 "part2-sample.txt"))
(println 5435 (part2 "part2-actual.txt"))
