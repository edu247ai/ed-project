from .model import AnomalyModel
from edu247model.ml.preprocess.clean import clean_df

def train_anomaly_from_dataset(df, column):
    df = clean_df(df)
    values = df[column].tolist()

    model = AnomalyModel()
    model.fit(values)
    return model
