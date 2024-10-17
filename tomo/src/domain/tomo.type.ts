// ToMo指数 = (楽しさ指数 × 10 + 目的指数 × 5 + 可能性指数 × 1.66) - (感情的圧力指数 × 1.66 + 経済的圧力指数 × 5 + 惰性指数 × 10)

/**
 * TOMOの回答となる指数
 * 1(まったく違う)~7(まったくその通りである)
 */
type Index = number;

// See https://www.factor.ai/survey
export type ToMoSurveyAnswer = {
  /** 楽しさ */
  play: Index;
  /** 目的 */
  purpose: Index;
  /** 可能性 */
  potential: Index;
  /** 感情的圧力 */
  emotionalPressure: Index;
  /** 経済的圧力 */
  economicPressure: Index;
  /** 惰性 */
  inertia: Index;
};

export type MotivationKey = keyof ToMoSurveyAnswer;
