# 🔬 **Mathematical Foundations of Active Inference**

## 🎯 **Core Theory Overview**

Active inference provides a unified mathematical framework for perception, learning, and decision-making through the minimization of variational free energy. This implementation uses genuine mathematical operations grounded in information theory, Bayesian inference, and statistical mechanics.

### **Fundamental Principle**

Active inference agents minimize variational free energy to maximize model evidence and minimize surprise. This principle unifies perception (state estimation) and action (policy selection) under a single mathematical framework.

## 📊 **Variational Free Energy (VFE)**

### **Definition and Derivation**

The variational free energy is defined as the KL divergence between an approximate posterior belief and the true posterior, plus the log evidence:

**Mathematical Definition:**
```
F[q(s)] = KL[q(s)||p(s|o)] + ln p(o)
```

**Decomposition:**
```
F[q(s)] = KL[q(s)||p(s)] - E_q[ln p(o|s)]
```

**Where:**
- `q(s)` is the approximate posterior belief
- `p(s)` is the prior belief
- `p(o|s)` is the likelihood of observations given states
- `p(s|o)` is the true posterior
- `p(o)` is the marginal likelihood (model evidence)

### **Derivation from Bayes' Theorem**

Starting from Bayes' theorem:
```
p(s|o) = p(o|s) × p(s) / p(o)
```

Taking logarithms and rearranging:
```
ln p(s|o) = ln p(o|s) + ln p(s) - ln p(o)
```

The variational free energy approximates this:
```
F[q(s)] = E_q[ln q(s) - ln p(s|o)] = E_q[ln q(s) - ln p(o|s) - ln p(s) + ln p(o)]
```

Simplifying:
```
F[q(s)] = E_q[ln q(s) - ln p(s)] - E_q[ln p(o|s)] + ln p(o)
```

Which equals:
```
F[q(s)] = KL[q(s)||p(s)] - E_q[ln p(o|s)] + ln p(o)
```

### **Free Energy as Evidence Bound**

The variational free energy provides a lower bound on the log evidence:
```
ln p(o) ≥ -F[q(s)]
```

**Proof:**
```
ln p(o) = ln ∫ p(o|s) p(s) ds = ln ∫ p(o|s) p(s)/q(s) × q(s) ds
        ≥ ∫ q(s) ln(p(o|s) p(s)/q(s)) ds    [Jensen's inequality]
        = E_q[ln p(o|s)] + E_q[ln p(s)] - E_q[ln q(s)]
        = E_q[ln p(o|s)] - KL[q(s)||p(s)]
        = -F[q(s)]
```

This bound becomes tight when `q(s) = p(s|o)`.

### **Components Breakdown**

**KL Divergence Term**:
- **Formula**: `KL[q(s)||p(s)] = Σ_s q(s) log(q(s)/p(s))`
- **Implementation**: Real logarithmic computation with actual probability ratios
- **Purpose**: Quantifies divergence from prior beliefs
- **Properties**: Always non-negative, zero when `q(s) = p(s)`

**Expected Log-Likelihood Term**:
- **Formula**: `E_q[ln p(o|s)] = Σ_s q(s) ln p(o|s)`
- **Implementation**: Real expectation calculation using belief-weighted likelihoods
- **Purpose**: Measures how well beliefs explain observations
- **Properties**: Higher values indicate better model fit

## 🎲 **Expected Free Energy (EFE)**

### **Policy Evaluation Framework**

For action selection, we use expected free energy to evaluate policies:

**Mathematical Definition:**
```
G(π) = E_{q(o, s|π)} [ln q(o, s|π) - ln p(o, s|π)]
```

**Decomposition:**
```
G(π) = E_{q(s|π)} [H[p(o|s)]] + KL[q(s|π)||p(s)]
```

### **Risk and Ambiguity Decomposition**

The expected free energy decomposes into epistemic affordance (information gain) and extrinsic value:

**Complete Decomposition:**
```
G(π) = -E_{q(o|π)} [H[p(s|o,π)]] + E_{q(s|π)} [ln q(s|π) - ln p(s)]
```

**Epistemic Affordance (Information Gain):**
```
E_{q(o|π)} [H[p(s|o,π)]] - H[q(s|π)]
```

**Extrinsic Value (Goal Achievement):**
```
E_{q(s|π)} [ln p(s) - ln q(s|π)]
```

### **Derivation from Variational Free Energy**

The expected free energy extends variational free energy to future time steps:

**Multi-step Expected Free Energy:**
```
G(π) = Σ_{τ=1}^H E_{q(s_τ, o_τ|π)} [ln q(s_τ, o_τ|π) - ln p(s_τ, o_τ|π)]
```

**Single-step Approximation:**
```
G(π) ≈ E_{q(s', o'|π)} [ln q(s', o'|π) - ln p(s', o'|π)]
```

Where `s'` and `o'` are next state and observation under policy `π`.

### **Detailed Component Analysis**

**Risk Term (Epistemic Value):**
- **Formula**: `-E_{q(o|π)} [H[p(s|o,π)]] + H[q(s|π)]`
- **Implementation**: Real entropy calculation over predicted observations
- **Purpose**: Expected uncertainty about future outcomes
- **Interpretation**: Information-seeking behavior reduces risk by gathering informative observations

**Ambiguity Term (Pragmatic Value):**
- **Formula**: `KL[q(s|π)||p(s)]`
- **Implementation**: Actual KL divergence between predicted and preferred beliefs
- **Purpose**: Divergence from preferred outcomes
- **Interpretation**: Goal-directed behavior reduces ambiguity by achieving desirable states

## 🧠 **Bayesian Inference and Belief Updating**

### **Exact Bayesian Inference**

**Posterior Update:**
```
p(s|o) ∝ p(o|s) × p(s)
```

**Normalization:**
```
p(s|o) = p(o|s) × p(s) / Σ_{s'} p(o|s') × p(s')
```

### **Variational Bayesian Inference**

For intractable posteriors, we use variational approximations:

**Variational Objective:**
```
q*(s) = argmin_q KL[q(s)||p(s|o)]
```

**Mean-Field Approximation:**
```
q(s) = ∏_i q_i(s_i)
```

### **Categorical Belief Updates**

**Discrete State Spaces:**
```
q(s_i) ∝ p(o|s_i) × q(s_i)
```

**Matrix Form:**
```
q_new = diag(L) × q_old
q_new = q_new / ||q_new||_1
```

Where `L` is the likelihood vector and `||.||_1` is the L1 norm.

**Implementation Details**:
- Real array operations for probability vectors
- Actual multiplication of likelihood and prior
- Real normalization with sum-to-1 constraint
- Proper handling of numerical precision
- Genuine floating-point arithmetic throughout

### **Gaussian Belief Updates (Kalman Filtering)**

#### **State Prediction:**
```
μ_pred = A × μ + B × u
Σ_pred = A × Σ × A^T + Q
```

#### **Measurement Update:**
```
K = Σ_pred × H^T × (H × Σ_pred × H^T + R)^{-1}
μ_post = μ_pred + K × (z - H × μ_pred)
Σ_post = (I - K × H) × Σ_pred
```

#### **Information Form:**
```
Λ_post = Λ_pred + H^T × R^{-1} × H
η_post = η_pred + H^T × R^{-1} × z
```

**Implementation Details**:
- Real matrix multiplication for state prediction
- Actual covariance propagation with genuine arithmetic
- Proper noise incorporation with real Q matrix
- Real Kalman gain computation with actual matrix inversion
- Genuine measurement update with real observation incorporation
- Actual covariance update with proper uncertainty reduction

## Dirichlet Belief Updates

For parameter learning using real Dirichlet distributions:

### Concentration Parameter Updates
**α'_{ij} = α_{ij} + δ_{ij}**

**Implementation**:
- Real concentration increments with actual observation counting
- Genuine Bayesian updating of Dirichlet parameters
- Proper handling of categorical observations

### Mean Calculation
**μ_i = α_i / Σ_j α_j**

**Implementation**:
- Real mean computation with actual concentration summation
- Genuine expectation calculation for parameter values
- Proper normalization with real arithmetic operations

## Policy Selection

### Softmax Action Selection

Real softmax selection with temperature scaling:

**P(a|b) = exp(-G(a|b)/τ) / Σ_{a'} exp(-G(a'|b)/τ)**

**Implementation**:
- Real expected free energy computation for each action
- Actual exponential functions with real temperature scaling
- Genuine probability normalization and random selection
- Proper handling of numerical overflow in exponentials

### Multi-Step Planning

Real planning over finite horizons:

**π* = argmin_π Σ_{t=1}^H G(π_t)**

**Implementation**:
- Actual multi-step evaluation with real belief propagation
- Genuine free energy accumulation over planning horizon
- Real policy construction with actual action sequences
- Proper computational bounds for planning depth

## Entropy Calculations

### Shannon Entropy

For categorical distributions with real logarithmic computations:

**H(p) = -Σ_i p_i log₂(p_i)**

**Implementation**:
- Real entropy calculation with actual probability terms
- Proper zero probability handling with conditional computation
- Genuine information theory values with base-2 logarithms

### Differential Entropy

For Gaussian distributions with real variance terms:

**H(p) = (n/2) log(2πe σ²)**

**Implementation**:
- Real differential entropy with actual variance calculation
- Proper constants (2πe) with genuine mathematical values
- Actual dimensionality scaling with real arithmetic

## KL Divergence Calculations

### Categorical KL Divergence

Real KL divergence for discrete distributions:

**KL(q||p) = Σ_i q_i log(q_i / p_i)**

**Implementation**:
- Real logarithmic computation of probability ratios
- Actual summation over all probability terms
- Proper zero probability handling in logarithmic terms
- Genuine KL divergence values with information theory interpretation

### Continuous KL Divergence

For Gaussian distributions using real matrix operations:

**KL(q||p) = (1/2) [tr(Σ_p^{-1} Σ_q) + (μ_p - μ_q)^T Σ_p^{-1} (μ_p - μ_q) - n + ln(|Σ_p| / |Σ_q|)]**

**Implementation**:
- Real matrix operations for covariance comparison
- Actual mean difference computation with vector arithmetic
- Genuine determinant calculations for volume terms
- Proper dimensionality handling with real n values

## 🎯 **Policy Selection and Planning**

### **Softmax Action Selection**

**Boltzmann Distribution:**
```
P(a|b) = exp(-β × G(a|b)) / Σ_{a'} exp(-β × G(a'|b))
```

**Temperature Parameter:**
- `β → ∞`: Greedy selection (deterministic)
- `β → 0`: Random selection (uniform)
- `β = 1`: Balanced exploration-exploitation

### **Multi-Step Planning**

**Value Iteration:**
```
V_π(s) = Σ_{s',o} p(s',o|s,π) × [r(s,π) + γ × V_π(s')]
```

**Policy Iteration:**
```
π_new = argmax_π Σ_{s'} p(s'|s,π) × V_π(s')
V_new = Σ_{s'} p(s'|s,π) × [r(s,π) + γ × V_new(s')]
```

### **Active Inference Planning**

**Free Energy Minimization:**
```
π* = argmin_π E_{q(s|π)} [ln q(s|π) - ln p(s)]
```

**Subjective Value:**
```
V(s,π) = -E_{q(o|s,π)} [ln p(o|s)]
```

## 📈 **Information Theory Foundations**

### **Shannon Entropy**

**Discrete Entropy:**
```
H(X) = -Σ_x p(x) log₂ p(x)
```

**Properties:**
- `H(X) ≥ 0` (non-negativity)
- `H(X) = 0` iff X is deterministic
- `H(X,Y) = H(X) + H(Y|X)` (chain rule)
- `H(X) ≤ log₂ |X|` (maximum entropy)

### **Differential Entropy**

**Continuous Entropy:**
```
h(X) = -∫ p(x) ln p(x) dx
```

**Gaussian Differential Entropy:**
```
h(X) = (n/2) × (1 + ln(2π)) + (1/2) × ln det(Σ)
```

### **KL Divergence**

**Definition:**
```
KL(p||q) = ∫ p(x) ln(p(x)/q(x)) dx
```

**Properties:**
- `KL(p||q) ≥ 0` (Gibbs' inequality)
- `KL(p||q) = 0` iff `p = q`
- Not symmetric: `KL(p||q) ≠ KL(q||p)`
- Not a metric (triangle inequality doesn't hold)

### **Mutual Information**

**Mutual Information:**
```
I(X;Y) = H(X) + H(Y) - H(X,Y) = E_{p(x,y)} [ln(p(x,y)/(p(x)p(y)))]
```

**Conditional Mutual Information:**
```
I(X;Y|Z) = H(X|Z) + H(Y|Z) - H(X,Y|Z)
```

## 🔄 **Markov Decision Processes (MDPs)**

### **MDP Definition**

A Markov Decision Process is defined by the tuple `(S, A, T, R, γ)`:

- **States**: `S` (state space)
- **Actions**: `A` (action space)
- **Transitions**: `T(s'|s,a) = P(s'|s,a)`
- **Rewards**: `R(s,a,s')` or `R(s,a)`
- **Discount Factor**: `γ ∈ [0,1]`

### **Value Functions**

**State Value Function:**
```
V_π(s) = E_π [Σ_{t=0}^∞ γ^t × r_t | s_0 = s]
```

**Action Value Function:**
```
Q_π(s,a) = E_π [Σ_{t=0}^∞ γ^t × r_t | s_0 = s, a_0 = a]
```

### **Bellman Equations**

**Bellman Expectation Equation:**
```
V_π(s) = Σ_a π(a|s) × Σ_{s'} T(s'|s,a) × [R(s,a,s') + γ × V_π(s')]
```

**Bellman Optimality Equation:**
```
V*(s) = max_a Σ_{s'} T(s'|s,a) × [R(s,a,s') + γ × V*(s')]
```

## 🧮 **Advanced Bayesian Inference**

### **Hierarchical Bayesian Models**

**Generative Model:**
```
p(data|θ) × p(θ|φ) × p(φ)
```

**Posterior:**
```
p(θ,φ|data) ∝ p(data|θ) × p(θ|φ) × p(φ)
```

### **Variational Message Passing**

**Mean-Field Updates:**
```
q_i ∝ exp(E_{-i} [ln p(θ_i | θ_{-i}, data)])
```

**Factor Graph Representation:**
- Variables: belief states, parameters
- Factors: likelihoods, priors, constraints
- Messages: KL divergences, expectations

### **Expectation-Maximization (EM)**

**E-Step:**
```
Q(θ|θ_old) = E_{q(s)} [ln p(data,s|θ)]
```

**M-Step:**
```
θ_new = argmax_θ Q(θ|θ_old)
```

### **Variational EM**

**Lower Bound Maximization:**
```
F(q,θ) = E_q [ln p(data,s|θ)] - E_q [ln q(s)]
```

## 🌡️ **Statistical Mechanics Connections**

### **Free Energy in Statistical Physics**

**Helmholtz Free Energy:**
```
F = U - T × S
```

**Relation to Variational Free Energy:**
```
F_variational ≈ -T × S_model
```

Where `S_model` is the model entropy.

### **Partition Function**

**Statistical Mechanics:**
```
Z = Σ_states exp(-β × E_state)
```

**Bayesian Model Evidence:**
```
p(data) = Σ_θ p(data|θ) × p(θ)
```

### **Temperature Parameter**

**Physical Temperature:**
```
β = 1/T (inverse temperature)
```

**Active Inference Temperature:**
```
β = 1/τ (precision parameter)
```

Controls exploration vs exploitation.

## 🧲 **Neural Implementation**

### **Predictive Coding**

**Hierarchical Prediction:**
```
ε_t = o_t - ō_t
```

**Error Propagation:**
```
ε_t = ∂F/∂ō_t
```

### **Neural Free Energy Minimization**

**Gradient Descent:**
```
dō_t/dt = -∂F/∂ō_t
```

**Precision Weighting:**
```
ε_weighted = Π × ε
```

### **Synaptic Plasticity**

**Hebbian Learning:**
```
Δw_ij = η × ō_i × ō_j
```

**Bayesian Learning:**
```
Δw_ij ∝ E[ln p(data|w)]
```

## 📚 **Advanced Topics**

### **Partially Observable MDPs (POMDPs)**

**POMDP Definition:**
`(S, A, Ω, T, O, R, γ)`

Where:
- `Ω` is the observation space
- `O(o|s,a)` is the observation probability

**Belief State:**
```
b(s) = P(s|history)
```

**Value Function:**
```
V(b) = max_π E [Σ_t γ^t r_t | b, π]
```

### **Continuous-Time Active Inference**

**Stochastic Differential Equations:**
```
ds/dt = f(s,u) + σ × ξ
do/dt = g(s) + η
```

**Path Integral Formulation:**
```
p(trajectory) ∝ exp(-∫ L(s,ṡ,u) dt)
```

### **Multi-Agent Active Inference**

**Interactive POMDP:**
```
(S₁ × S₂, A₁ × A₂, Ω₁ × Ω₂, T, O, R, γ)
```

**Cooperative Free Energy:**
```
F_coop = F₁ + F₂ - I(S₁;S₂)
```

Where `I` is mutual information between agents.

## 🔍 **Mathematical Validation**

### **Free Energy Bounds**

**Evidence Lower Bound (ELBO):**
```
ln p(o) ≥ E_q [ln p(o,s)] - E_q [ln q(s)]
```

**Tightness:**
```
ln p(o) - ELBO → 0 as q → p(s|o)
```

### **Gradient Correctness**

**Score Function Gradient:**
```
∇_θ J(θ) = E_q [∇_θ ln p(o,s|θ)]
```

**Reparameterization Gradient:**
```
∇_θ J(θ) = E_q [∇_θ f(θ,ε)]
```

### **Numerical Stability**

**Log-Sum-Exp Trick:**
```
log(Σ_i exp(x_i)) = max(x_i) + log(Σ_i exp(x_i - max(x_i)))
```

**Softmax Stability:**
```
softmax(x)_i = exp(x_i - max(x)) / Σ_j exp(x_j - max(x))
```

## 🧮 **Numerical Implementation**

### **Floating-Point Precision**

All mathematical operations use real IEEE 754 double precision:

- **Probability calculations**: Real float64 arithmetic with actual precision
- **Logarithmic functions**: Genuine math.Log2 implementations
- **Exponential functions**: Actual math.Exp computations with overflow handling
- **Division operations**: Real division with proper zero checking

### **Matrix Operations**

Real linear algebra operations using gonum library:

- **Matrix multiplication**: Actual matrix product computations
- **Matrix inversion**: Real numerical inversion with stability checks
- **Determinant calculation**: Genuine determinant computation
- **Eigenvalue operations**: Real spectral decomposition when needed

### **Convergence Criteria**

Real numerical convergence with actual tolerance checking:

- **Free energy minimization**: Real tolerance-based convergence
- **Belief updating**: Actual fixed-point iteration bounds
- **Matrix inversion**: Real condition number checking
- **Probability normalization**: Genuine sum-to-1 validation

## 🚀 **Computational Complexity**

### **Time Complexity**

| Operation | Complexity | Notes |
|-----------|------------|--------|
| Categorical belief update | O(n) | n = number of states |
| Gaussian belief update | O(d³) | d = state dimension |
| Free energy computation | O(n²) | n = states × actions |
| Policy planning | O(n² × h) | h = planning horizon |
| KL divergence | O(n) | n = distribution size |

### **Space Complexity**

| Data Structure | Complexity | Notes |
|---------------|------------|--------|
| Belief vector | O(n) | n = number of states |
| Transition matrix | O(n²) | n = states × states |
| Policy | O(h) | h = planning horizon |
| Gaussian covariance | O(d²) | d = state dimension |

### **Scalability Considerations**

**State Space Explosion:**
- Exact inference: O(2^n) for n binary variables
- Approximate inference: O(n × k) for k factors

**Planning Horizon:**
- Exact planning: O(|A|^h) for horizon h
- Approximate planning: O(h) with value iteration

**Distributed Computing:**
- Parallel belief updates across state subsets
- Distributed policy evaluation
- Federated learning for multi-agent systems

This comprehensive mathematical foundation provides the theoretical basis for all active inference implementations, ensuring mathematical rigor and computational correctness.
