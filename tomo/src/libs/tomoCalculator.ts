import type { ToMoSurveyAnswer } from "./tomo.type";

/**
 * ToMo指数を計算する。
 *
 * 小数点第1位で四捨五入し、整数を返す。
 * @see https://qiita.com/quechon167/items/d92220290a11bfafe074
 */
export const calculateToMo = (answer: ToMoSurveyAnswer): number => {
  const direct =
    answer.play * 10 + answer.purpose * 5 + answer.potential * 1.66;
  const indirect =
    answer.emotionalPressure * 1.66 +
    answer.economicPressure * 5 +
    answer.inertia * 10;
  return Math.round(direct - indirect);
};
