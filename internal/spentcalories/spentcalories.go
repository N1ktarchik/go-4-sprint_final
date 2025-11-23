package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	mas := strings.Split(data, ",")

	if len(mas) != 3 {
		err := errors.New("data parse error")
		return 0, "", time.Duration(0), err
	}

	steps, err := strconv.Atoi(mas[0])

	if err != nil {
		return 0, "", time.Duration(0), err
	}

	if steps <= 0 {
		err := errors.New("the number of steps is negative or equal to zero")
		return 0, "", time.Duration(0), err
	}

	durationOfActivity, err := time.ParseDuration(mas[2])

	if err != nil {
		return 0, "", time.Duration(0), err
	}

	if durationOfActivity <= time.Duration(0) {
		err := errors.New("the duration of the activity is too short")
		return 0, "", time.Duration(0), err
	}

	return steps, mas[1], durationOfActivity, nil
}

func distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient

	return float64(steps) * stepLength / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {

	if duration <= time.Duration(0) {
		return 0.0
	}

	dist := distance(steps, height)

	return dist / duration.Hours()

}

func TrainingInfo(data string, weight, height float64) (string, error) {

	steps, typeOfActivity, duration, err := parseTraining(data)

	if err != nil {
		log.Println(err)
	}

	switch typeOfActivity {

	case "Бег":

		dist := distance(steps, height)

		averageSpeed := meanSpeed(steps, height, duration)

		calorise, err := RunningSpentCalories(steps, weight, height, duration)

		if err != nil {
			return "", err
		}

		result := fmt.Sprintf("Тип тренировки: Бег\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			duration.Hours(), dist, averageSpeed, calorise)

		return result, nil

	case "Ходьба":

		dist := distance(steps, height)

		averageSpeed := meanSpeed(steps, height, duration)

		calorise, err := WalkingSpentCalories(steps, weight, height, duration)

		if err != nil {
			return "", err
		}

		result := fmt.Sprintf("Тип тренировки: Ходьба\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			duration.Hours(), dist, averageSpeed, calorise)

		return result, nil

	default:

		err := errors.New("неизвестный тип тренировки")
		return "", err
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {
		err := errors.New("the number of steps is negative or equal to zero")
		return 0.0, err
	}

	if weight <= 0 {
		err := errors.New("the user's weight is negative or zero")
		return 0.0, err
	}

	if height <= 0 {
		err := errors.New("the user's height is negative or zero")
		return 0.0, err
	}

	if duration <= time.Duration(0) {
		err := errors.New("The running time is too short")
		return 0.0, err
	}

	averageSpeed := meanSpeed(steps, height, duration)

	calories := weight * averageSpeed * duration.Minutes() / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {
		err := errors.New("the number of steps is negative or equal to zero")
		return 0.0, err
	}

	if weight <= 0 {
		err := errors.New("the user's weight is negative or zero")
		return 0.0, err
	}

	if height <= 0 {
		err := errors.New("the user's height is negative or zero")
		return 0.0, err
	}

	if duration <= time.Duration(0) {
		err := errors.New("The walking time is too short")
		return 0.0, err
	}

	averageSpeed := meanSpeed(steps, height, duration)

	calories := weight * averageSpeed * duration.Minutes() / minInH

	return calories * walkingCaloriesCoefficient, nil
}
