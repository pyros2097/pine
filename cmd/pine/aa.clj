(struct user
  :name str required min(10) max(10) regex("[a-z]")  
  :year int required min(0) max(150))

(defn user/fullname [u] 
  (`Mr. ${u.name}`))