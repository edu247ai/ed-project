from .env_autoscale import AutoscaleEnv
from .agent_ppo import PPOAgent
from .memory import ReplayBuffer

def train_rl_episodes(episodes: int = 100):
    env = AutoscaleEnv()
    agent = PPOAgent()
    memory = ReplayBuffer()

    for _ in range(episodes):
        state = env.reset()
        done = False

        while not done:
            action = agent.act(state)
            next_state, reward, done = env.step(action)
            memory.add(state, action, reward, next_state, done)
            state = next_state

    agent.update(memory.all())
    return agent
