(enum direction
  :left
  :right)

(enum position
  :top
  :bottom
  :left
  :right)

(struct user
  :name str required min(10) max(10) regex("[a-z]")
  :year int required min(0) max(150))

(defn user/fullname [u] 
  (`Mr. ${u.name}`))

(def state {})

(defn use-state [initial]
  # (def key (get :componentIndex))
  (map/getset state key initial))

(defn counter [:title str]
  (var index (use-state 10))
  (var text (title + "test" + (str index))
  (var u1 (user/new :name "test" :year 123))
  (defn plus []
    (set index (index + 1)))
  (defn minus []
    (set index (index - 1)))
  
  (html/div :column :align-center :justify-center margin(10)
   (html/div "hello")
   (html/div (user/fullname u1))
   (html/div (user/name u1))
   (html/div
    (html/div text)
    (html/button :click plus  "+")
    (html/button :click minus "-")))

(defn show-dir [:dir direction]
  (html/div :dir))

(defn app []
  (html/div
    (counter :title "My Counter")))

(var items array/new 100)
(each range 10
  (array/push items i))

(test "for-equal-or-not"
  (assert 1 == 2)
  (assert 2 == 14))
