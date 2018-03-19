from sklearn.base import BaseEstimator, TransformerMixin
import numpy as np

class ColumnSelector(BaseEstimator, TransformerMixin):

    def __init__(self, columns=None):
        self.columns = columns

    def fit(self, X, y=None):
        return self

    def transform(self, X, y=None):
        return X[self.columns]
    

class ColumnImputer(BaseEstimator, TransformerMixin):
    """Impute missing values.

    Columns of dtype object are imputed with the most frequent value 
    in column.

    Columns of other types are imputed with mean of column.

    Credits: http://stackoverflow.com/questions/25239958/impute-categorical-missing-values-in-scikit-learn
    """

    def __init__(self, strategy="mean"):
        self.strategy = strategy

    def fit(self, X, y=None):
        if X.dtype == np.dtype("O"):
            X = X.astype("str")
        # Get Fill Counts and store the frequent value
        values, counts = np.unique(X.astype("str"), return_counts=True)
        self.fill = values[np.argmax(counts)]
        
        return self

    def transform(self, X, y=None):
        if X.dtype == np.dtype("O"):
            nan_idxs = np.where(X.astype("str") == 'nan')
        else:
            nan_idxs = np.where(np.isnan(X))
        X[nan_idxs] = self.fill
        return X