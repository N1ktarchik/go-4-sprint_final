package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	mas := strings.Split(data, ",")

	if len(mas) != 2 {
		err := errors.New("data parse error")
		return 0, 0, err
	}

	steps, err := strconv.Atoi(mas[0])

	if err != nil {
		return 0, 0, err
	}

	if steps <= 0 {
		err := errors.New("the number of steps is less than or equal to 0.")
		return 0, 0, err
	}

	walkDuration, err := time.ParseDuration(mas[1])

	if err != nil {
		return 0, 0, err
	}

	if walkDuration <= time.Duration(0) {
		err := errors.New("the duration of the activity is too short")
		return 0, 0, err

	}

	return steps, walkDuration, nil

}

func DayActionInfo(data string, weight, height float64) string {

	steps, walkDuration, err := parsePackage(data)

	if err != nil {
		log.Print(err)
		return ""
	}

	distance := float64(steps) * stepLength / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, walkDuration)

	if err != nil {
		log.Print(err)
		return ""
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)

	return result

}
