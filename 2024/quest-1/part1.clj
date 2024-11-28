(def sample-input "ABBAC")
(def potion-costs {\A 0 \B 1 \C 3})
(defn potion-cost [input]
  (reduce + (map potion-costs input)))

(println 5 (potion-cost sample-input))

(def actual-input "ACABBBCBBCAABCCCCBCACABCCBCABBBAAABCAABBAAACBACACBCBACBCBCCBCAACABCCBABBACAAAABCBCCCBCBBCCACCACACBACBCBCCBCACBCBBBACCABCCCCBACABCACCABAABABACACACCAAAACCAABABABACBCBCBAACBBCCABAABABCABCABAABCCBBCACCCBAABCCBCCAACABCCCABABBBCBCCBBABABCCCBBBCACCCAAAABBBBACBACBAABCBACAACCCBCCBABBBBCABBCACCCBAACCBCCBACBAABCCCBBAACCCAACBABCBBBBCABCAABBBAACBACABBBCABBBACCBBCACCABBBBBABCABACBBACABAABACCCCACBBBBCCABBAACBCBABCCCAABCBABACCCAABAACBABACCBBACACCCCBBCAABABCABABAABBBACBBCCBCCCBCBCACCBCAACACCACAAACAAACCACBBCABABBABACBCABBCACBABBCACAABBCAACBACAACCBCBCCCBCBCABBBCACBBCABABABABCABABACACCACBCACBABAAACBABAAABBCBCABBBABCBBCACCBABBABCBBBBACCBCBBBBCACCABBBCACACCCBACACBCABBBCCAACCACBCBBCAAABCACCABBBBACCABBCAACBCCBAACABCACAAAAAAAABBBBCBCABBCCCAABCAACACABABCCABBCBCCCCACBCABABCABCABBCCBCACCABCCBACCBCCBCBBABCCCCACBABAAAAACBCCAABBBBABBBBCCBBCAAAACBBBCBBBCBCBCBBCBACBACACACCCABCAAACAAABACBABCAAAAACCCCBBCACABBBBCBCCBBACCCCCCCCBAACCABBBABBCAAABBBBACBCCCCCCBACABCACCABBCBABABBBBACACBCCBBAACBABAACAACCABAAABCBCCBBABCCCBABCBAC")

(println 1395 (potion-cost actual-input))
