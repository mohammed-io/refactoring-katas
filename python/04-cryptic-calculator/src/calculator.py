class Calculator:
    def __init__(self):
        pass

    def normalize(self, scores, new_min, new_max):
        normalized = []
        current_min = float("inf")
        current_max = float("-inf")

        for score in scores:
            if score < current_min:
                current_min = score
            if score > current_max:
                current_max = score

        score_range = current_max - current_min
        new_range = new_max - new_min

        for score in scores:
            try:
                normalized_score = new_min + ((score - current_min) / score_range) * new_range
            except ZeroDivisionError:
                normalized_score = float("nan")
            normalized.append(round(normalized_score, 2))

        return normalized
