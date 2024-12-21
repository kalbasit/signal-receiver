{ self, ... }:
{
  perSystem =
    {
      lib,
      pkgs,
      ...
    }:
    {
      packages.signal-api-receiver =
        let
          shortRev = self.shortRev or self.dirtyShortRev;
        in
        pkgs.buildGoModule {
          name = "signal-api-receiver-${shortRev}";

          src = lib.fileset.toSource {
            fileset = lib.fileset.unions [
              ../../cmd
              ../../go.mod
              ../../go.sum
              ../../receiver
              ../../server
            ];
            root = ../..;
          };

          subPackages = [ "./cmd/signal-api-receiver" ];

          vendorHash = "sha256-0Qxw+MUYVgzgWB8vi3HBYtVXSq/btfh4ZfV/m1chNrA=";

          doCheck = true;

          meta = {
            description = "Signal API receiver";
            homepage = "https://github.com/kalbasit/signal-api-receiver";
            license = lib.licenses.mit;
            maintainers = [ lib.maintainers.kalbasit ];
          };
        };

    };
}
