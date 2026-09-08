from .env_autoscale import AutoscaleEnv
from .agent_ppo import PPOAgent

def evaluate_agent(agent: PPOAgent, episodes: int = 10):
    env = AutoscaleEnv()
    results = []

    for _ in range(episodes):
        state = env.reset()
        done = False
        total_reward = 0

        while not done:
            action = agent.act(state)
            next_state, reward, done = env.step(action)
            total_reward += reward
            state = next_state

        results.append(total_reward)

    return results
