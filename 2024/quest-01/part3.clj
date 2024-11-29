(def potion-costs {\A 0 \B 1 \C 3 \D 5})
(defn clean-group [group] (remove #(= % \x) group))
(defn solo-cost [solo] (reduce + (map potion-costs solo)))
(defn group-cost [group]
  (case (count group)
    1 (+ 0 (solo-cost group))
    2 (+ 2 (solo-cost group))
    3 (+ 6 (solo-cost group))
    0))

(defn potion-cost [input]
  (->> input
       (partition 3)
       (map clean-group)
       (map group-cost)
       (reduce +)))

(println 1     (potion-cost "xBx"))
(println 6     (potion-cost "AAA"))
(println 15    (potion-cost "BCD"))
(println 8     (potion-cost "xCC"))
(println 30    (potion-cost (slurp "part3-sample.txt")))
(println 27723 (potion-cost (slurp "part3-actual.txt")))
