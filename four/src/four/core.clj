(ns four.core
	(:gen-class))

(require 'digest)

(defn do-until [f x p]
  (if (p x) x (recur f (f x) p)))

(defn int-md5 [x]
	(digest/md5 (str "yzbqklnj" x)))

(defn -main [& args]
  (println "Welcome to my project! These are your args:" args)
  (do-until inc 0 #(= (subs (int-md5 %) 0 5) "00000")))