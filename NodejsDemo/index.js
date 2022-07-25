console.log('Loading function');

exports.handler = async (event, context, callback) => {
    var adder1 = 4;
    var adder2 = 9;
    var randomNum = Math.floor(Math.random());
    var ans = adder1 * randomNum + adder2 * (1 - randomNum);
    callback(null, ans);
};
