TURNSTILE = ( passenger -> TURNSTILE ).
const N = 3
LOOP = LOOP[0],
LOOP[n: 0..N] = (
    when n < N passenger -> LOOP[n+1]
  | when n == N depart -> LOOP[n]
  // A special action to skip directly to 0 from any state
  | arrive -> LOOP[0]
).
CAR = ( depart -> ride -> arrive -> CAR ).
||ROLLER_COASTER = ( TURNSTILE || LOOP || CAR ).
