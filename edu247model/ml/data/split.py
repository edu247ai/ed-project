def train_test_split(df, ratio=0.8):
    n = int(len(df) * ratio)
    return df[:n], df[n:]
