import random

def random_walk(k, max_steps):
    # Starting position at the center
    x, y = 0, 0
    

    directions = [('up', (0, 1)), ('down', (0, -1)), ('right', (1, 0)), ('left', (-1, 0))]
    
    steps_taken = 0
    #max_steps = k**2
    while steps_taken < max_steps:
        direction = random.choice(directions)
        x += direction[1][0]
        y += direction[1][1]
        

        if x > k or x < -k or y > k or y < -k:

            return steps_taken, (x, y)
        
        steps_taken += 1
    

    return steps_taken, (x, y)

k = 5 
max_steps = 100 
random_walk_result = random_walk(k, max_steps)
print(random_walk_result)
