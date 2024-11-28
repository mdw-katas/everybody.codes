(def sample-input 
  (str/join "\n"
   ["WORDS:THE,OWE,MES,ROD,HER"
    ""
    "AWAKEN THE POWER, ADORNED WITH THE FLAMES BRIGHT IRE"]))

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

(println (part-1 sample-input))

(def actual-input
  (str/join "\n"
    ["WORDS:LOR,LL,SI,OR,CU,EX,DO"
     ""
     "LOREM IPSUM DOLOR SIT AMET, CONSECTETUR ADIPISCING ELIT, SED DO EIUSMOD TEMPOR INCIDIDUNT UT LABORE ET DOLORE MAGNA ALIQUA. UT ENIM AD MINIM VENIAM, QUIS NOSTRUD EXERCITATION ULLAMCO LABORIS NISI UT ALIQUIP EX EA COMMODO CONSEQUAT. DUIS AUTE IRURE DOLOR IN REPREHENDERIT IN VOLUPTATE VELIT ESSE CILLUM DOLORE EU FUGIAT NULLA PARIATUR. EXCEPTEUR SINT OCCAECAT CUPIDATAT NON PROIDENT, SUNT IN CULPA QUI OFFICIA DESERUNT MOLLIT ANIM ID EST LABORUM."]))

(println (part-1 actual-input))