{
  imports = [
    ./docker.nix
    ./signal-api-receiver.nix
  ];

  perSystem =
    { config, ... }:
    {
      packages.default = config.packages.signal-api-receiver;
    };
}
