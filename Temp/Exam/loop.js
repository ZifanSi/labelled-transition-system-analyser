// Maximum value for the counter
const MAX = 4

// Define a loop process with conditions and actions
LOOP = LOOP[0], 
LOOP[n: 0..MAX] = (
    // Increment the counter if it is less than MAX
    when n < MAX increment -> LOOP[n+1]
    // Reset the counter to 0 when it reaches MAX
  | when n == MAX reset -> LOOP[0]
    // A special action to skip directly to 0 from any state
  | skip -> LOOP[0]
).
