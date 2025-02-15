# scheduling

## abstract concepts 

**sleep block**

- a certain amount of time spent sleeping, decision-making capacity is restored over this time.

**motivation**

- motivation is something that compels you to do something. you decide to do the thing that has the highest motivation value.
- motivation can be derived from various sources:
   - prospect of dopamine:
      - how likely you are to yield dopamine from a task influences your decision on whether or not to begin on it.
      - sources of dopamine:
         - success: "a feeling of a job well done" yields dopamine.
            - if you are "succeeding" already at the task at hand, you are more likely to continue doing the task at hand since you have increased confidence of success, which correlates with a likely source of dopamine. this is also known as "flow".
         - super-stimuli: socializing, entertainment, drugs, gambling. these yield an extreme amount of dopamine relative to other tasks.
            - engaging in these will heavily increase the likelihood of you choosing to continue to do these in the future since you can yield a massive amount of dopamine relatively confidently.
   - willpower:
      - a conscious, rational decision to work on a task, this is like a muscle, while it will get tired in the short-term, the total amount it can lift can be increased in the long term. it can also restore energy in the short-term.

this scheduler is concerned with how to maximize your motivation over the long-term for working (and succeeding) at the tasks you define.

the "natural" motivation value of a task is simply the sum of all the sources excluding willpower.

the question then becomes: *what tasks should you spend willpower on, and how much willpower to spend.*

**task-switching**

- people have a limited task-switching resource, switching tasks decreases concentration
- concentration is like the efficacy of your willpower (consider a river, the amount of water in the river is constant, that is your willpower capacity, the speed the water is moving at is your concentration)

**decision-making capacity**

- people have a limited capacity to make decisions, the more decisions you make, the more *decision fatigue* builds up.
- *decision fatigue* causes poorer choices over time. (since you think less before making a decision)

**emotional regulation**

- people have a limited capacity to mitigate the effects of emotions on their person.
   - ex. a person who spends some emotional regulation capacity on staying calm in a meeting will have a harder time staying calm given further stressful situations.

## realized concepts

**decision-making capacity**

- decision-making capacity is defined as minutes of work on a task with low/medium/high decision-making per cycle.

**willpower**

- willpower is one's short-term capacity to make up for a lack of natural motivation to do something specific (instead of its alternatives which may have higher natural motivation).
- we will assume that all tasks have a uniform natural motivation: normal by default (unless explicitly specified to be high) (natural motivation will be considered to be low if decision-making required is high)
- frequent task-switching will decrease willpower.
- willpower can be expended over the short-term, but may also be replenished over the short-term.

**task template**

- optional preset fields for a certain type of task.
- estimated time, decision-making difficulty, deadline

**task**

- estimated time required: (some amount of minutes)
- decision-making difficulty: (low, medium, high)
   - measures the level of decision-making involved in this task.
- time till deadline: (some amount of minutes)

**scheduling strategy**

1. earliest deadlines come first.
2. effort balancing. (avoid stacking low natural motivation tasks back to back)
3. tasks with most decision-making required should come first.

