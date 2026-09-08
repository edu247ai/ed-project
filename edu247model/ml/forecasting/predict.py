from .model import ForecastModel

def forecast_series(model: ForecastModel, series, steps=1):
    return model.forecast(series, steps)
