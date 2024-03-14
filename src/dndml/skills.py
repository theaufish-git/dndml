from enum import Enum, auto

class Skills(Enum):
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

    def __init__(self, value):
        if isinstance(value, str):
            value = Skills.from_str(value).value

        print(value)
        super().__init__(value)


    @staticmethod
    def from_str(s: str):
        match s.lower():
            case "acrobatics":
                return Skills.ACROBATICS
            case "animal_handling":
                return Skills.ANIMAL_HANDLING
            case "arcana":
                return Skills.ARCANA
            case "athletics":
                return Skills.ATHLETICS
            case "deception":
                return Skills.DECEPTION
            case "history":
                return Skills.HISTORY
            case "insight":
                return Skills.INSIGHT
            case "intimidation":
                return Skills.INTIMIDATION
            case "investigation":
                return Skills.INVESTIGATION
            case "medicine":
                return Skills.MEDICINE
            case "nature":
                return Skills.NATURE
            case "perception":
                return Skills.PERCEPTION
            case "performance":
                return Skills.PERFORMANCE
            case "persuasion":
                return Skills.PERSUASION
            case "religion":
                return Skills.RELIGION
            case "sleight_of_hand":
                return Skills.SLEIGHT_OF_HAND
            case "stealth":
                return Skills.STEALTH
            case "survival":
                return Skills.SURVIVAL
