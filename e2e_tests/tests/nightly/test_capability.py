import pytest

from tests import config as conf
from tests import experiment as exp


@pytest.mark.nightly  # type: ignore
def test_mmdetection_pytorch_const() -> None:
    config = conf.load_config(conf.cv_examples_path("mmdetection_pytorch/const_fake_data.yaml"))
    config = conf.set_max_length(config, {"batches": 200})

    exp.run_basic_test_with_temp_config(config, conf.cv_examples_path("mmdetection_pytorch"), 1)


@pytest.mark.nightly  # type: ignore
def test_bert_glue_const() -> None:
    config = conf.load_config(conf.nlp_examples_path("bert_glue_pytorch/const.yaml"))
    config = conf.set_max_length(config, {"batches": 200})

    exp.run_basic_test_with_temp_config(config, conf.nlp_examples_path("bert_glue_pytorch"), 1)


@pytest.mark.nightly  # type: ignore
def test_gaea_pytorch_const() -> None:
    config = conf.load_config(conf.nas_examples_path("gaea_pytorch/eval/const.yaml"))
    config = conf.set_max_length(config, {"batches": 200})

    exp.run_basic_test_with_temp_config(config, conf.nas_examples_path("gaea_pytorch/eval"), 1)


@pytest.mark.nightly  # type: ignore
def test_gan_mnist_pytorch_const() -> None:
    config = conf.load_config(conf.gan_examples_path("gan_mnist_pytorch/const.yaml"))
    config = conf.set_max_length(config, {"batches": 200})

    exp.run_basic_test_with_temp_config(config, conf.gan_examples_path("gan_mnist_pytorch"), 1)


@pytest.mark.nightly  # type: ignore
def test_detr_coco_pytorch_const() -> None:
    config = conf.load_config(conf.cv_examples_path("detr_coco_pytorch/const_fake.yaml"))
    config = conf.set_max_length(config, {"batches": 200})

    exp.run_basic_test_with_temp_config(config, conf.cv_examples_path("detr_coco_pytorch"), 1)


@pytest.mark.nightly  # type: ignore
def test_deformabledetr_coco_pytorch_const() -> None:
    config = conf.load_config(conf.cv_examples_path("deformabledetr_coco_pytorch/const_fake.yaml"))
    config = conf.set_max_length(config, {"batches": 200})

    exp.run_basic_test_with_temp_config(
        config, conf.cv_examples_path("deformabledetr_coco_pytorch"), 1
    )


@pytest.mark.nightly  # type: ignore
def test_supclr_pytorch_const() -> None:
    config = conf.load_config(conf.cv_examples_path("contrastive_learning_pytorch/supclr.yaml"))
    config = conf.set_global_batch_size(config, 16)
    config = conf.set_slots_per_trial(config, 1)
    config = conf.set_max_length(config, {"batches": 200})

    exp.run_basic_test_with_temp_config(
        config, conf.cv_examples_path("contrastive_learning_pytorch"), 1
    )


@pytest.mark.nightly  # type: ignore
def test_simclr_pytorch_const() -> None:
    config = conf.load_config(conf.cv_examples_path("contrastive_learning_pytorch/simclr.yaml"))
    config = conf.set_global_batch_size(config, 16)
    config = conf.set_slots_per_trial(config, 1)
    config = conf.set_max_length(config, {"batches": 200})

    exp.run_basic_test_with_temp_config(
        config, conf.cv_examples_path("contrastive_learning_pytorch"), 1
    )


@pytest.mark.nightly  # type: ignore
def test_supclr_train_eval_pytorch_const() -> None:
    config = conf.load_config(
        conf.cv_examples_path("contrastive_learning_pytorch/supclr_train_eval.yaml")
    )
    config = conf.set_global_batch_size(config, 16)
    config = conf.set_slots_per_trial(config, 1)
    config = conf.set_max_length(config, {"batches": 200})

    exp.run_basic_test_with_temp_config(
        config, conf.cv_examples_path("contrastive_learning_pytorch"), 1
    )


@pytest.mark.nightly  # type: ignore
def test_simclr_train_eval_pytorch_const() -> None:
    config = conf.load_config(
        conf.cv_examples_path("contrastive_learning_pytorch/simclr_train_eval.yaml")
    )
    config = conf.set_global_batch_size(config, 16)
    config = conf.set_slots_per_trial(config, 1)
    config = conf.set_max_length(config, {"batches": 200})

    exp.run_basic_test_with_temp_config(
        config, conf.cv_examples_path("contrastive_learning_pytorch"), 1
    )
