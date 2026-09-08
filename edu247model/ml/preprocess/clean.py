def clean_df(df):
    df = df.dropna()
    df = df.drop_duplicates()
    return df
