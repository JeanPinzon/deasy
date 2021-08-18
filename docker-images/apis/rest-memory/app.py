import yaml

from py_easy_rest.server import App


def _read_config_from_yaml():
    with open(r"./api-definition.yaml") as file:
        return yaml.load(file, Loader=yaml.FullLoader)


if __name__ == '__main__':
    config = _read_config_from_yaml()
    pyrApp = App(config)

    pyrApp.app.run(
        host='0.0.0.0',
        port=8000,
        debug=True,
        auto_reload=True,
    )