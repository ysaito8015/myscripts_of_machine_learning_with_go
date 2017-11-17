import pandas as pd

# 列名の定義
cols = [
        'integercolumn',
        'stringcolumn'
        ]

# pandas を利用してCSVを読み込む
data = pd.read_csv('myfile.csv', names=cols)

# integer column の最大値を表示する
print(data['integercolumn'].max())
