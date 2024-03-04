const { fromRpcSig } = require("ethereumjs-util");
const { signTypedData } = require("@metamask/eth-sig-util");

const signPermit = (domain, message, privateKey) => {
  const { name, version, chainId, verifyingContract } = domain;

  const EIP712Domain = [
    { name: "name", type: "string" },
    { name: "version", type: "string" },
    { name: "chainId", type: "uint256" },
    { name: "verifyingContract", type: "address" },
  ];

  const Permit = [
    { name: "owner", type: "address" },
    { name: "spender", type: "address" },
    { name: "value", type: "uint256" },
    { name: "nonce", type: "uint256" },
    { name: "deadline", type: "uint256" },
  ];

  const data = {
    primaryType: "Permit",
    types: { EIP712Domain, Permit },
    domain: { name, version, chainId, verifyingContract },
    message: message,
  };

  return fromRpcSig(signTypedData({ privateKey, data, version: "V4" }));
};

module.exports = {
  signPermit,
};
