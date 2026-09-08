def normalize_column(df, col):
    mn = df[col].min()
    mx = df[col].max()
    df[col] = (df[col] - mn) / (mx - mn)
    return df
