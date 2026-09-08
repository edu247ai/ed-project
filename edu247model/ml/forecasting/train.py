from .model import ForecastModel
from edu247model.ml.preprocess.clean import clean_df

def train_forecast_from_dataset(df, column):
    df = clean_df(df)
    series = df[column].tolist()
    model = ForecastModel(window=5)
    return model, series
