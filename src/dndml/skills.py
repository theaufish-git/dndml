from .enum import StrEnum, auto
from .stats import Stats


class Skills(StrEnum):
    ACROBATICS = auto()
    ANIMAL_HANDLING = auto()
    ARCANA = auto()
    ATHLETICS = auto()
    DECEPTION = auto()
    HISTORY = auto()
    INSIGHT = auto()
    INTIMIDATION = auto()
    INVESTIGATION = auto()
    MEDICINE = auto()
    NATURE = auto()
    PERCEPTION = auto()
    PERFORMANCE = auto()
    PERSUASION = auto()
    RELIGION = auto()
    SLEIGHT_OF_HAND = auto()
    STEALTH = auto()
    SURVIVAL = auto()

    def stat(self):
        return {
            Skills.ACROBATICS: Stats.DEX,
            Skills.ANIMAL_HANDLING: Stats.WIS,
            Skills.ARCANA: Stats.INT,
            Skills.ATHLETICS: Stats.STR,
            Skills.DECEPTION: Stats.CHA,
            Skills.HISTORY: Stats.INT,
            Skills.INSIGHT: Stats.WIS,
            Skills.INTIMIDATION: Stats.CHA,
            Skills.INVESTIGATION: Stats.INT,
            Skills.MEDICINE: Stats.WIS,
            Skills.NATURE: Stats.INT,
            Skills.PERCEPTION: Stats.WIS,
            Skills.PERFORMANCE: Stats.CHA,
            Skills.PERSUASION: Stats.CHA,
            Skills.RELIGION: Stats.INT,
            Skills.SLEIGHT_OF_HAND: Stats.DEX,
            Skills.STEALTH: Stats.DEX,
            Skills.SURVIVAL: Stats.WIS,
        }[self]
