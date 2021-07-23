from copy import deepcopy
from dataclasses import dataclass, field
from typing import List

import fire
import numpy as np
from dataclasses_json import dataclass_json
from numpy.random import randint


@dataclass_json
@dataclass
class Frames:
    items: List["Frame"] = field(default_factory=list)


@dataclass_json()
@dataclass
class Frame:
    id: int = 0
    medium_id: int = 0
    block_campaign_categories: List[int] = field(default_factory=list)
    block_creative_tags: List[int] = field(default_factory=list)


def gen(seed: int = 1, samples: int = 100):
    np.random.seed(seed)
    frames1 = Frames()
    frames1.items = []
    for i in range(samples):
        frame = Frame()
        frame.id = i + 1
        frame.medium_id = int(randint(65535))
        frame.block_campaign_categories = list(map(int, randint(65535, size=randint(100))))
        frame.block_creative_tags = list(map(int, randint(65535, size=randint(100))))
        frames1.items.append(frame)

    frames2 = deepcopy(frames1)
    for i in range(0, samples, 2):
        frame = Frame()
        frame.id = i + 1
        frame.medium_id = int(randint(65535))
        frame.block_campaign_categories = list(map(int, randint(65535, size=randint(100))))
        frame.block_creative_tags = list(map(int, randint(65535, size=randint(100))))
        frames2.items[i] = frame

    with open("frames_1.json", "w") as f:
        f.write(frames1.to_json())

    with open("frames_2.json", "w") as f:
        f.write(frames2.to_json())


    print(frames1.to_json())


if __name__ == '__main__':
    fire.Fire(gen)
