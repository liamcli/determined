package searcher

import (
	"testing"

	"gotest.tools/assert"

	"github.com/determined-ai/determined/master/pkg/nprand"
	"github.com/determined-ai/determined/master/pkg/schemas/expconf"
)

func TestNestedHyperparameters(t *testing.T) {
	spec := expconf.Hyperparameters{
		"optimizer": expconf.Hyperparameters{
			"type": expconf.Hyperparameter{
				RawConstHyperparameter: &expconf.ConstHyperparameter{RawVal: "adam"},
			},
			"learning_rate": expconf.Hyperparameter{
				RawConstHyperparameter: &expconf.ConstHyperparameter{RawVal: 0.01},
			},
			"momentum": expconf.Hyperparameter{
				RawConstHyperparameter: &expconf.ConstHyperparameter{RawVal: 0.9},
			},
		},
	}
	hps := spec["optimizer"].(expconf.Hyperparameters)
	assert.Equal(
		t,
		hps["type"].(expconf.Hyperparameter).RawConstHyperparameter.RawVal,
		"adam",
	)
	assert.Equal(
		t,
		hps["learning_rate"].(expconf.Hyperparameter).RawConstHyperparameter.RawVal,
		0.01,
	)
	assert.Equal(
		t,
		hps["momentum"].(expconf.Hyperparameter).RawConstHyperparameter.RawVal,
		0.9,
	)

	flatSpec := expconf.Hyperparameters{
		"optimizer.type": expconf.Hyperparameter{
			RawConstHyperparameter: &expconf.ConstHyperparameter{RawVal: "adam"},
		},
		"optimizer.learning_rate": expconf.Hyperparameter{
			RawConstHyperparameter: &expconf.ConstHyperparameter{RawVal: 0.01},
		},
		"optimizer.momentum": expconf.Hyperparameter{
			RawConstHyperparameter: &expconf.ConstHyperparameter{RawVal: 0.9},
		},
	}
	flatHPs := expconf.FlattenHPs(spec)
	assert.DeepEqual(t, flatHPs, flatSpec)

	nestedHPs := expconf.UnflattenHPs(flatHPs)
	hps = nestedHPs["optimizer"].(expconf.Hyperparameters)
	assert.Equal(
		t,
		hps["type"].(expconf.Hyperparameter).RawConstHyperparameter.RawVal,
		"adam",
	)
	assert.Equal(
		t,
		hps["learning_rate"].(expconf.Hyperparameter).RawConstHyperparameter.RawVal,
		0.01,
	)
	assert.Equal(
		t,
		hps["momentum"].(expconf.Hyperparameter).RawConstHyperparameter.RawVal,
		0.9,
	)
}

func TestNestedSampling(t *testing.T) {
	spec := expconf.Hyperparameters{
		"optimizer": expconf.Hyperparameters{
			"type": expconf.Hyperparameter{
				RawConstHyperparameter: &expconf.ConstHyperparameter{RawVal: "adam"},
			},
			"learning_rate": expconf.Hyperparameter{
				RawConstHyperparameter: &expconf.ConstHyperparameter{RawVal: 0.01},
			},
			"momentum": expconf.Hyperparameter{
				RawConstHyperparameter: &expconf.ConstHyperparameter{RawVal: 0.9},
			},
		},
	}
	flatSpec := expconf.FlattenHPs(spec)
	rand := nprand.New(0)
	sample := sampleAll(spec, rand)
	flatSample := sampleAll(flatSpec, rand)
	nestedSample := unflattenSample(flatSample)
	assert.DeepEqual(t, sample, nestedSample)
}

func TestSamplingReproducibility(t *testing.T) {
	spec := expconf.Hyperparameters{
		"cat": expconf.Hyperparameter{
			RawCategoricalHyperparameter: &expconf.CategoricalHyperparameter{
				RawVals: []interface{}{0, 1, 2, 3, 4, 5, 6},
			},
		},
		"const": expconf.Hyperparameter{
			RawConstHyperparameter: &expconf.ConstHyperparameter{RawVal: "val"},
		},
		"double": expconf.Hyperparameter{
			RawDoubleHyperparameter: &expconf.DoubleHyperparameter{RawMinval: 0, RawMaxval: 100},
		},
		"int": expconf.Hyperparameter{
			RawIntHyperparameter: &expconf.IntHyperparameter{RawMinval: 0, RawMaxval: 100},
		},
		"log": expconf.Hyperparameter{
			RawLogHyperparameter: &expconf.LogHyperparameter{RawBase: 10, RawMinval: -2, RawMaxval: 2},
		},
	}

	// Run the checks multiple times; map iteration order, if it has an effect on the result, is an
	// uncontrollable source of randomness that may end up allowing the checks to pass incorrectly.
	for seed := uint32(0); seed < 50; seed++ {
		// Check that sampling twice gives the same result both times.
		rand1 := nprand.New(seed)
		rand2 := nprand.New(seed)

		sample1 := sampleAll(spec, rand1)
		sample2 := sampleAll(spec, rand2)

		assert.Equal(t, 5, len(sample1))
		assert.Equal(t, 5, len(sample2))

		for name, val1 := range sample1 {
			val2, ok := sample2[name]
			assert.Assert(t, ok)
			assert.Equal(t, val1, val2)
		}

		// Compare direct RNG output to check by proxy whether the internal states have stayed the same.
		assert.Equal(t, rand1.Bits64(), rand2.Bits64())
	}
}
