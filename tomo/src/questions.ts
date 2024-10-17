import { ToMoSurveyAnswer } from "./domain/tomo.type";

export type ToMoQuestion = {
  readonly [key in keyof ToMoSurveyAnswer]: {
    readonly text: string;
  };
};

export const tomoQuestions: ToMoQuestion = {
  play: { text: "今の仕事を続けているのは仕事そのものが楽しいから" },
  purpose: {
    text: "今の仕事を続けているのはその仕事に重要な目的があると思うから",
  },
  potential: {
    text: "今の仕事を続けているのは自分の目標を達成する上で有益だから",
  },
  emotionalPressure: {
    text: "今の仕事を続けているのは辞めたら自分と自分のことを気にかけてくれる人を落胆させてしまうから",
  },
  economicPressure: {
    text: "今の仕事を続けているのはこの仕事を失ったら金銭上の目標を達成できなくなるから",
  },
  inertia: { text: "今の仕事を続ける妥当な理由はない" },
} as const;
