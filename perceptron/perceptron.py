import random
import pandas as pd


def generate_random(start, end, dim, num):
    data = []
    for x in range(num):
        x1 = random.uniform(start, end)
        x2 = random.uniform(start, end)
        data.append((x1, x2))
    df = pd.DataFrame(data)
    df.columns = ["x1", "x2"]
    df["x0"] = 1
    df = df[["x0", "x1", "x2"]]
    return df


def assign_target(df):
    targets = []

    for i, row in df.iterrows():
        if (row["x1"] + 2 * row["x2"]) > 0:
            targets.append(0)
        else:
            targets.append(1)

    print(len(targets))
    df["y"] = targets
    return df


def get_random_weights(size):
    targets = []
    for i in range(size):
        targets.append(random.uniform(0, 1))
    return targets


def main():
    df = generate_random(-5, 5, 2, 1000)
    df_with_target = assign_target(df)
    weights = get_random_weights(2)

    print(df_with_target)
    print(df_with_target["y"].value_counts())
    print(weights)


if __name__ == "__main__":
    main()
