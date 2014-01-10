(define xx 2)
(define y 8)
(define z 7)

(describe arithmetic-test
    (== (+ 5 5) 10)
    (== (- 10 7) 3)
    (== (* 2 4)  8)
    (== (/ 25 5) 5)
    (== (% 7 5)  2))

(describe arithmetic-limits
          (== (- 5 9) 0))

(describe condition-test
    (== (< xx y)  #t)
    (== (< y z)  #f)
    (== (> xx y)  #f)
    (== (> z xx)  #t)
    (== (<= xx 2) #t)
    (== (>= z 7) #t)
    (== (!= xx xx) #f)
    (== (!= 2 xx) #f)
    (== (!= 2 3) #t)
    (== (! #f)   #t)
    (== (! #t)   #f)
    (== (or #f #f)     #f)
    (== (or #f #f #t)  #t)
    (== (and #t #f #t) #f)
    (== (and #t #t #t) #t)
    (== (or (> xx z) (> y z))  #t)
    (== (and (> xx z) (>y z))  #f)
    (== (even? 2) #t)
    (== (even? 3) #f)
    (== (odd? 3) #t)
    (== (odd? 2) #f))
