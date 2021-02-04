use wasplib::client::*;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call("your_sc_request", your_sc_request);
    exports.add_view("your_sc_view", your_sc_view);
}

fn your_sc_request(ctx: &ScCallContext) {
    ctx.log("your_sc_request");
}

fn your_sc_view(ctx: &ScViewContext) {
    ctx.log("your_sc_view");
}