# Implemented method {#sec:method}

For a state indexed by $i$, with supplied belief mass $q_i$ and matching
observation likelihood $L_i$, the discrete update weights the prior by $L_i$
and normalizes by the total weight when it is positive. An absent observation
row for the supplied action contributes unit likelihood. An absent observation
entry within a present row contributes zero weight.

The engine reports a variational free-energy score as divergence from the
uniform model prior minus its expected log-likelihood term. Its expected
free-energy score adds an observation-entropy term and a policy divergence.
The rollout scores actions greedily at each horizon step and predicts the next
belief from the transition table. The score implementation follows the local
code contract. More complete formulations of active inference are described in
[@parr2022active].
