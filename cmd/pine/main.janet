(use jaylib)

(def [width height] [800 450])
(var [x y radius] [(/ width 2) (/ height 2) 50])


(defn main [& args]
  (init-window width height "Jay")
  (set-target-fps 60)

  (while (not (window-should-close))
    (when (key-pressed? :q)
      (close-window))

    (when (key-down? :right)
      (if (< x 300)
        (set x (+ x 3))))
    (when (key-down? :left)
      (if (> x 0)
        (set x (- x 3))))
    (when (key-down? :down)
      (set y (+ y 3)))
    (when (key-down? :up)
      (set y (- y 3)))
    
    (begin-drawing)
    (clear-background [22 22 0])
     
    (draw-text "move ball" 10 10 20 :white)
    (draw-circle (+ x radius) (+ y radius) radius :maroon)
    (draw-fps 50 50)
    (end-drawing))

  (close-window))

